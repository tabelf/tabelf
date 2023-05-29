package audit

import (
	"context"
	entstation "tabelf/backend/gen/entschema/station"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuditGoodStationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAuditGoodStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuditGoodStationLogic {
	return &GetAuditGoodStationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuditGoodStationLogic) GetAuditGoodStation(req *types.GetAuditGoodStationRequest) (resp *types.GetAuditGoodStationResponse, err error) {
	hasAuthority, err := base.HasAdminAuthority(l.ctx, req.UserUID)
	if err != nil {
		return nil, err
	}
	if !hasAuthority {
		return nil, app.ErrAccountAdminNotExist(l.ctx)
	}
	stations, err := app.EntClient.Station.Query().Where(
		entstation.Status(app.Hidden),
		entstation.DeactivatedAtIsNil(),
	).All(l.ctx)
	if err != nil {
		return nil, err
	}
	data := make([]*types.GoodStationRecommend, 0)
	for _, station := range stations {
		data = append(data, &types.GoodStationRecommend{
			UID:         station.UID,
			Title:       station.Title,
			Description: station.Description,
			Image:       station.Image,
			Tags:        station.Tags,
			Icon:        station.Icon,
			Source:      station.Source,
			Link:        station.Link,
		})
	}
	return &types.GetAuditGoodStationResponse{
		Data: data,
	}, nil
}
