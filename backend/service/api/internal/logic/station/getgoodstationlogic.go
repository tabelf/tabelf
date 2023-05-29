package station

import (
	"context"
	entstation "tabelf/backend/gen/entschema/station"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodStationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodStationLogic {
	return &GetGoodStationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodStationLogic) GetGoodStation(req *types.GetGoodStationRequest) (resp *types.GetGoodStationResponse, err error) {
	hasAuthority, err := base.HasAdminAuthority(l.ctx, req.UserUID)
	if err != nil {
		return nil, err
	}
	if !hasAuthority {
		return nil, app.ErrAccountAdminNotExist(l.ctx)
	}
	station, err := app.EntClient.Station.Query().Where(
		entstation.UID(req.StationUID),
		entstation.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		return nil, err
	}
	return &types.GetGoodStationResponse{
		StationUID:  station.UID,
		CategoryUID: station.CategoryUID,
		SiteName:    station.Source,
		Link:        station.Link,
		Title:       station.Title,
		Description: station.Description,
		Image:       station.Image,
		Tags:        station.Tags,
	}, nil
}
