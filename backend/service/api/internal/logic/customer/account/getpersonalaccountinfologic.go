package account

import (
	"context"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPersonalAccountInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPersonalAccountInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPersonalAccountInfoLogic {
	return &GetPersonalAccountInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPersonalAccountInfoLogic) GetPersonalAccountInfo(req *types.GetPersonalAccountInfoRequest) (resp *types.GetPersonalAccountInfoResponse, err error) {
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
	return &types.GetPersonalAccountInfoResponse{
		UserUID:       account.UID,
		UserName:      account.Nickname,
		Email:         account.Email,
		Industry:      account.Industry,
		Image:         account.Image,
		Description:   account.Description,
		HasMembership: base.MembershipValidity(account),
		HasEntire:     account.HasEntire,
	}, nil
}
