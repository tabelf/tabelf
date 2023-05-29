package audit

import (
	"context"
	entstation "tabelf/backend/gen/entschema/station"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"
	"time"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAuditGoodStationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAuditGoodStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAuditGoodStationLogic {
	return &UpdateAuditGoodStationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAuditGoodStationLogic) UpdateAuditGoodStation(req *types.UpdateAuditGoodStationRequest) (resp *types.UpdateAuditGoodStationResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	hasAuthority, err := base.HasAdminAuthority(l.ctx, req.UserUID)
	if err != nil {
		return nil, err
	}
	if !hasAuthority {
		return nil, app.ErrAccountAdminNotExist(l.ctx)
	}

	stationUpdate := app.EntClient.Station.Update().SetStatus(req.Status)
	if !req.Status {
		stationUpdate.SetDeactivatedAt(time.Now())
	}
	save, err := stationUpdate.Where(
		entstation.UID(req.StationUID),
		entstation.DeactivatedAtIsNil(),
	).Save(l.ctx)
	if err != nil {
		return nil, err
	}
	if save == 0 {
		return nil, app.ErrStationAudit(l.ctx)
	}
	return &types.UpdateAuditGoodStationResponse{
		Message: app.HttpOK,
	}, nil
}
