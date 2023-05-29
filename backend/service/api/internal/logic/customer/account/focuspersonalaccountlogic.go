package account

import (
	"context"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	entfocus "tabelf/backend/gen/entschema/focus"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FocusPersonalAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFocusPersonalAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FocusPersonalAccountLogic {
	return &FocusPersonalAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FocusPersonalAccountLogic) FocusPersonalAccount(req *types.FocusPersonalAccountRequest) (resp *types.FocusPersonalAccountResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	focus, err := app.EntClient.Focus.Query().Where(
		entfocus.FollowerUID(req.UserUID),
		entfocus.FolloweeUID(req.FolloweeUID),
		entfocus.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil && !entschema.IsNotFound(err) {
		return nil, err
	}
	if err = app.WithTx(l.ctx, app.EntClient, func(tx *entschema.Tx) (err error) {
		if focus != nil {
			err = tx.Focus.Update().SetStatus(req.Status).
				Where(entfocus.ID(focus.ID)).
				Exec(l.ctx)
		} else {
			err = tx.Focus.Create().SetStatus(req.Status).
				SetFollowerUID(req.UserUID).
				SetFolloweeUID(req.FolloweeUID).
				Exec(l.ctx)
		}
		if err != nil {
			return err
		}
		num := app.Then(req.Status, 1, -1).(int)
		// 我的关注 + 1 / - 1
		if err = tx.Account.Update().AddFocus(num).Where(
			entaccount.UID(req.UserUID),
		).Exec(l.ctx); err != nil {
			return err
		}
		// ta的粉丝 + 1 / - 1
		if err = tx.Account.Update().AddFans(num).Where(
			entaccount.UID(req.FolloweeUID),
		).Exec(l.ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &types.FocusPersonalAccountResponse{
		Message: app.HttpOK,
	}, nil
}
