package community

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"strings"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"
	"time"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePublicCommunityImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePublicCommunityImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePublicCommunityImageLogic {
	return &UpdatePublicCommunityImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePublicCommunityImageLogic) UpdatePublicCommunityImage(req *types.UpdatePublicCommunityImageRequest) (resp *types.UpdatePublicCommunityImageResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}

	account, err := app.EntClient.Account.Query().Where(
		entaccount.UID(req.UserUID),
		entaccount.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		if entschema.IsNotFound(err) {
			return nil, app.ErrAccountInvalid(l.ctx)
		}
		return nil, err
	}

	exist, err := app.EntClient.PersonalFolder.Query().Where(
		entpersonalfolder.UID(req.FolderUID),
		entpersonalfolder.HasOpen(true),
		entpersonalfolder.DeactivatedAtIsNil(),
	).Exist(l.ctx)
	if err != nil {
		return nil, err
	}
	if !exist {
		return &types.UpdatePublicCommunityImageResponse{
			Message: app.HttpNO,
		}, nil
	}

	source := strings.TrimPrefix(req.Image, "data:image/png;base64,")
	data, err := base64.StdEncoding.DecodeString(source)
	if err != nil {
		return nil, app.ErrCommunityImageParse(l.ctx)
	}

	u, _ := url.Parse(app.TxCosCommunityURL)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		//设置超时时间
		Timeout: 10 * time.Second,
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:  app.Cos.SecretID,
			SecretKey: app.Cos.SecretKey,
		},
	})

	imageName := fmt.Sprintf("%s-%s.png", account.UID, req.FolderUID)
	if _, err = c.Object.Put(l.ctx, imageName, bytes.NewReader(data), nil); err != nil {
		return nil, err
	}
	return &types.UpdatePublicCommunityImageResponse{
		Message: app.HttpOK,
	}, nil
}
