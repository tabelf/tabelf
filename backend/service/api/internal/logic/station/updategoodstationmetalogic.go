package station

import (
	"context"
	"tabelf/backend/service/api/internal/logic/base"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGoodStationMetaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGoodStationMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGoodStationMetaLogic {
	return &UpdateGoodStationMetaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGoodStationMetaLogic) UpdateGoodStationMeta(req *types.UpdateGoodStationMetaRequest) (resp *types.UpdateGoodStationMetaResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	return base.UpdateGoodStationMeta(l.ctx, req.StationUID, req.UserUID, req.MetaType)
}
