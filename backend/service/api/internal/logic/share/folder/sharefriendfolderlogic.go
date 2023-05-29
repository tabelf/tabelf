package folder

import (
	"context"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"regexp"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"
	entsharelink "tabelf/backend/gen/entschema/sharelink"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareFriendFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareFriendFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareFriendFolderLogic {
	return &ShareFriendFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareFriendFolderLogic) ShareFriendFolder(req *types.ShareFriendFolderRequest) (resp *types.ShareFriendFolderResponse, err error) {
	if match, err := regexp.MatchString(app.EmailValidate, req.Email); err != nil || !match {
		return nil, app.ErrAccountEmailInvalid(l.ctx)
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
	shareLink, err := app.EntClient.ShareLink.Query().Where(
		entsharelink.UID(req.ShareUID),
		entsharelink.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		if entschema.IsNotFound(err) {
			return nil, app.ErrPersonalFolderShareLinkNotExist(l.ctx)
		}
		return nil, err
	}
	folder, err := app.EntClient.PersonalFolder.Query().Where(
		entpersonalfolder.UID(shareLink.FolderUID),
		entpersonalfolder.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		if entschema.IsNotFound(err) {
			return nil, app.ErrPersonalFolderNotExist(l.ctx)
		}
		return nil, err
	}
	emailURL := account.Email
	if app.IsNotBlank(emailURL) {
		emailURL = fmt.Sprintf("&nbsp;<a href=\"\" class=\"link\">%s</a>", emailURL)
	}
	e := &email.Email{
		To:      []string{req.Email},
		From:    app.Email.Username,
		Subject: app.EmailShareSubject,
		HTML: []byte(fmt.Sprintf(app.ShareEmailTemplate,
			req.Email,
			account.Nickname+emailURL,
			folder.FolderName,
			fmt.Sprintf("%s/v/%s", app.Basic.Domain, req.Authority+shareLink.UID),
			fmt.Sprintf("%s/v/%s", app.Basic.Domain, req.Authority+shareLink.UID),
		)),
	}
	if err = e.Send(app.Email.Addr+":"+app.EmailSSLPorts, smtp.PlainAuth(
		"",
		app.Email.Username,
		app.Email.Password,
		app.Email.Addr,
	)); err != nil {
		return nil, err
	}
	return &types.ShareFriendFolderResponse{
		Message: app.HttpOK,
	}, nil
}
