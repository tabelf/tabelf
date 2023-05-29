package station

import (
	"context"
	"tabelf/backend/gen/entschema"
	entstation "tabelf/backend/gen/entschema/station"
	entstationcategory "tabelf/backend/gen/entschema/stationcategory"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodStationCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodStationCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodStationCategoryLogic {
	return &GetGoodStationCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodStationCategoryLogic) GetGoodStationCategory(req *types.GetGoodStationCategoryRequest) (resp *types.GetGoodStationCategoryResponse, err error) {
	stationCategories, err := app.EntClient.StationCategory.Query().Where(
		entstationcategory.DeactivatedAtIsNil(),
		entstationcategory.Status(app.Show),
	).Order(entschema.Asc(entstationcategory.FieldSequence)).
		All(l.ctx)
	if err != nil {
		return nil, err
	}
	categories := make([]*types.GoodStationCategory, 0)
	for _, c := range stationCategories {
		categories = append(categories, &types.GoodStationCategory{
			UID:  c.UID,
			Name: c.Name,
		})
	}
	hasAdmin := false
	wait := 0
	if app.IsNotBlank(req.UserUID) {
		hasAdmin, err = base.HasAdminAuthority(l.ctx, req.UserUID)
		if err != nil {
			return nil, err
		}
		if hasAdmin {
			wait, err = app.EntClient.Station.Query().Where(
				entstation.Status(app.Hidden),
				entstation.DeactivatedAtIsNil(),
			).Count(l.ctx)
			if err != nil {
				return nil, err
			}
		}
	}
	return &types.GetGoodStationCategoryResponse{
		GoodStationCategories: categories,
		HasAuthority:          hasAdmin,
		Wait:                  wait,
	}, nil
}
