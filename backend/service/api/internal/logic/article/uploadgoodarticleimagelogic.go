package article

import (
	"bytes"
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"
	"time"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadGoodArticleImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadGoodArticleImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadGoodArticleImageLogic {
	return &UploadGoodArticleImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadGoodArticleImageLogic) UploadGoodArticleImage(req *types.UploadGoodArticleImageRequest) (resp *types.UploadGoodArticleImageResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	if len(req.File) == 0 {
		return nil, app.ErrCustomerHeaderImageEmpty(l.ctx)
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

	u, _ := url.Parse(app.TxCosArticleURL)
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

	imageName := fmt.Sprintf("%s-%s-%s", account.UID, app.GetYMD(time.Now()), req.Filename)
	if _, err = c.Object.Put(l.ctx, imageName, bytes.NewReader(req.File), nil); err != nil {
		return nil, err
	}
	return &types.UploadGoodArticleImageResponse{
		ImageURL: app.TxCosArticleURL + "/" + imageName,
	}, nil
}
