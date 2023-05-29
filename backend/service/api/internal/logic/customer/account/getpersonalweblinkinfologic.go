package account

import (
	"context"
	"github.com/shopspring/decimal"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type GetPersonalWebLinkInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPersonalWebLinkInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPersonalWebLinkInfoLogic {
	return &GetPersonalWebLinkInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPersonalWebLinkInfoLogic) GetPersonalWebLinkInfo(req *types.GetPersonalWebLinkInfoRequest) (resp *types.GetPersonalWebLinkInfoResponse, err error) {
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
	if app.StringSliceContains([]string{
		app.MonthMemberUser,
		app.YearMemberUser,
	}, account.MemberType) && time.Now().Before(account.MemberExpired) {
		return &types.GetPersonalWebLinkInfoResponse{
			UsedQuantity:  account.URLCount,
			TotalQuantity: account.URLLimit,
			Percent:       100,
			HasLimit:      false,
		}, nil
	}
	percent := int(app.DivDecimal(account.URLCount, account.URLLimit).
		RoundFloor(2).
		Mul(decimal.NewFromInt(100)).
		IntPart())
	if account.URLCount > account.URLLimit {
		percent = 100
	}
	return &types.GetPersonalWebLinkInfoResponse{
		UsedQuantity:  account.URLCount,
		TotalQuantity: account.URLLimit,
		Percent:       percent,
		HasLimit:      true,
	}, nil
}
