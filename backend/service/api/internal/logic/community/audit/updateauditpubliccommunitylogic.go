package audit

import (
	"context"
	entcommunity "tabelf/backend/gen/entschema/community"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAuditPublicCommunityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAuditPublicCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAuditPublicCommunityLogic {
	return &UpdateAuditPublicCommunityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAuditPublicCommunityLogic) UpdateAuditPublicCommunity(req *types.UpdateAuditPublicCommunityRequest) (resp *types.UpdateAuditPublicCommunityResponse, err error) {
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
	save, err := app.EntClient.Community.Update().SetStatus(req.Status).
		Where(
			entcommunity.UID(req.CommunityUID),
			entcommunity.DeactivatedAtIsNil(),
		).Save(l.ctx)
	if err != nil {
		return nil, err
	}
	if save == 0 {
		return nil, app.ErrCommunityAudit(l.ctx)
	}
	return &types.UpdateAuditPublicCommunityResponse{
		Status: req.Status,
	}, nil
}
