package account

import (
	"context"
	"regexp"
	"tabelf/backend/service/api/models"

	"github.com/zeromicro/go-zero/core/logx"

	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type UpdatePersonalAccountInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePersonalAccountInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePersonalAccountInfoLogic {
	return &UpdatePersonalAccountInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePersonalAccountInfoLogic) UpdatePersonalAccountInfo(req *types.UpdatePersonalAccountInfoRequest) (resp *types.UpdatePersonalAccountInfoResponse, err error) {
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
	if app.IsNotBlank(req.Email) {
		if match, err := regexp.MatchString(app.EmailValidate, req.Email); err != nil || !match {
			return nil, app.ErrAccountEmailInvalid(l.ctx)
		}
		if err = app.EntClient.Account.Update().
			SetEmail(req.Email).
			Where(entaccount.ID(account.ID)).
			Exec(l.ctx); err != nil {
			return nil, err
		}
	}
	if app.IsNotBlank(req.Description) {
		if req.Description == "-1" {
			req.Description = ""
		}
		if err = app.EntClient.Account.Update().
			SetDescription(req.Description).
			Where(entaccount.ID(account.ID)).
			Exec(l.ctx); err != nil {
			return nil, err
		}
	}
	if app.IsNotBlank(req.Industry) {
		if req.Industry == "-1" {
			req.Industry = ""
		}
		if err = app.EntClient.Account.Update().
			SetIndustry(req.Industry).
			Where(entaccount.ID(account.ID)).
			Exec(l.ctx); err != nil {
			return nil, err
		}
	}
	if app.IsNotBlank(req.UserName) {
		if err = app.EntClient.Account.Update().
			SetNickname(req.UserName).
			Where(entaccount.ID(account.ID)).
			Exec(l.ctx); err != nil {
			return nil, err
		}
	}
	if !account.HasEntire {
		account, err = app.EntClient.Account.Query().Where(
			entaccount.ID(account.ID),
		).First(l.ctx)
		if err != nil {
			return nil, err
		}
		if app.IsNotBlank(account.Nickname) &&
			app.IsNotBlank(account.Industry) &&
			app.IsNotBlank(account.Email) &&
			app.IsNotBlank(account.Description) {
			// 完善基本信息奖励
			if err = app.EntClient.Account.Update().
				AddURLLimit(app.InviteRewardFileNum).
				SetHasEntire(true).
				Where(entaccount.ID(account.ID)).
				Exec(l.ctx); err != nil {
				return nil, err
			}
			// 发送消息进行通知
			if err := models.SendMessage(
				l.ctx,
				app.SystemMessage,
				app.SystemAccount,
				account.UID,
				app.PerfectUserBaseInfoMessageTemplate,
			); err != nil {
				return nil, err
			}
		}
	}
	return &types.UpdatePersonalAccountInfoResponse{
		Message: app.HttpOK,
	}, nil
}
