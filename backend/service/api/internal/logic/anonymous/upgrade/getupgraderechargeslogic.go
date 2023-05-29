package upgrade

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	entrecharge "tabelf/backend/gen/entschema/recharge"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type GetUpgradeRechargesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUpgradeRechargesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeRechargesLogic {
	return &GetUpgradeRechargesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUpgradeRechargesLogic) GetUpgradeRecharges(req *types.GetUpgradeRechargesRequest) (resp *types.GetUpgradeRechargesResponse, err error) {
	recharges, err := app.EntClient.Recharge.Query().Where(
		entrecharge.DeactivatedAtIsNil(),
	).All(l.ctx)
	if err != nil {
		return nil, err
	}
	upgrades := make([]*types.UpgradeRecharges, 0)
	defaultAmount := ""
	for _, recharge := range recharges {
		if recharge.Default {
			defaultAmount = recharge.Amount
		}
		upgrades = append(upgrades, &types.UpgradeRecharges{
			UID:          recharge.UID,
			Title:        recharge.Title,
			OriginAmount: recharge.OriginAmount,
			Amount:       recharge.Amount,
			Descriptions: recharge.Descriptions,
			ThemeColor:   recharge.ThemeColor,
		})
	}
	return &types.GetUpgradeRechargesResponse{
		DefaultAmount:    defaultAmount,
		UpgradeRecharges: upgrades,
	}, nil
}
