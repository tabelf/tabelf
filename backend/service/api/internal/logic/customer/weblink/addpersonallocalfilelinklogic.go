package weblink

import (
	"bytes"
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"golang.org/x/sync/errgroup"
	"net/http"
	"net/url"
	"strings"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	entweblink "tabelf/backend/gen/entschema/weblink"
	entworkspace "tabelf/backend/gen/entschema/workspace"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"
	"time"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPersonalLocalFileLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPersonalLocalFileLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPersonalLocalFileLinkLogic {
	return &AddPersonalLocalFileLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPersonalLocalFileLinkLogic) AddPersonalLocalFileLink(req *types.AddPersonalLocalFileLinkRequest) (resp *types.AddPersonalLocalFileLinkResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	if len(req.File) == 0 {
		return nil, app.ErrCustomerLocalFileEmpty(l.ctx)
	}
	fileType := "zip"
	switch req.FileType {
	case "application/pdf":
		fileType = "pdf"
	case "image/jpeg", "image/png", "image/webp", "image/bmp", "image/gif":
		fileType = "jpg"
	case "text/csv":
		fileType = "csv"
	case "application/vnd.ms-excel", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":
		fileType = "xlsx"
	case "application/msword", "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
		fileType = "docx"
	case "video/mp4":
		fileType = "mp4"
	case "audio/mpeg":
		fileType = "mp3"
	case "application/vnd.ms-powerpoint", "application/vnd.openxmlformats-officedocument.presentationml.template":
		fileType = "ppt"
	default:
	}

	var (
		g         errgroup.Group
		account   *entschema.Account
		workspace *entschema.Workspace
		count     int
	)
	g.Go(func() (err error) {
		account, err = app.EntClient.Account.Query().Where(
			entaccount.UID(req.UserUID),
			entaccount.DeactivatedAtIsNil(),
		).First(l.ctx)
		if err != nil {
			if entschema.IsNotFound(err) {
				return app.ErrAccountInvalid(l.ctx)
			}
			return err
		}
		// 会员不做数量校验
		//if base.MembershipValidity(account) {
		//	return nil
		//}
		//if account.URLCount+1 > account.URLLimit {
		//	return app.ErrCustomerWebLinkLimit(l.ctx)
		//}
		return nil
	})
	g.Go(func() (err error) {
		count, err = app.EntClient.WebLink.Query().Where(
			entweblink.WorkspaceUID(req.WorkspaceUID),
			entweblink.DeactivatedAtIsNil(),
		).Count(l.ctx)
		return err
	})
	g.Go(func() (err error) {
		workspace, err = app.EntClient.Workspace.Query().Where(
			entworkspace.UID(req.WorkspaceUID),
			entworkspace.DeactivatedAtIsNil(),
		).First(l.ctx)
		return err
	})
	if err = g.Wait(); err != nil {
		return nil, err
	}

	u, _ := url.Parse(app.TxCosLocalFileURL)
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

	fileName := fmt.Sprintf("%s-%s-%s", account.UID, app.GetYMD(time.Now()), req.Filename)
	if _, err = c.Object.Put(l.ctx, fileName, bytes.NewReader(req.File), nil); err != nil {
		return nil, err
	}

	if err = app.WithTx(l.ctx, app.EntClient, func(tx *entschema.Tx) (err error) {
		if err = tx.WebLink.Create().
			SetTitle(req.Filename).
			SetDescription(strings.ToUpper(fileType)).
			SetImage("").
			SetLink(app.TxCosLocalFileURL + "/" + fileName).
			SetSequence(count).
			SetFileType(fileType).
			SetUserUID(req.UserUID).
			SetWorkspaceUID(req.WorkspaceUID).
			SetFolderUID(workspace.PersonalFolderUID).
			Exec(l.ctx); err != nil {
			return err
		}
		if err = tx.Account.Update().AddURLCount(1).
			Where(entaccount.UID(req.UserUID)).
			Exec(l.ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &types.AddPersonalLocalFileLinkResponse{
		Message: app.HttpOK,
	}, nil
}
