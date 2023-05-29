package audit

import (
	"context"
	"github.com/jinzhu/copier"
	entcommunity "tabelf/backend/gen/entschema/community"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuditPublicCommunityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAuditPublicCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuditPublicCommunityLogic {
	return &GetAuditPublicCommunityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuditPublicCommunityLogic) GetAuditPublicCommunity(req *types.GetAuditPublicCommunityRequest) (resp *types.GetAuditPublicCommunityResponse, err error) {
	hasAuthority, err := base.HasAdminAuthority(l.ctx, req.UserUID)
	if err != nil {
		return nil, err
	}
	if !hasAuthority {
		return nil, app.ErrAccountAdminNotExist(l.ctx)
	}
	count, err := app.EntClient.Community.Query().Where(
		entcommunity.Status(app.WaitAuditStatus),
		entcommunity.DeactivatedAtIsNil(),
	).Count(l.ctx)
	if err != nil {
		return nil, err
	}
	resp = &types.GetAuditPublicCommunityResponse{
		AuditCount:         count,
		PersonalWorkspaces: make([]*types.AuditWorkspace, 0),
	}
	if count != 0 {
		community, err := app.EntClient.Community.Query().Where(
			entcommunity.Status(app.WaitAuditStatus),
			entcommunity.DeactivatedAtIsNil(),
		).First(l.ctx)
		if err != nil {
			return nil, err
		}
		folder, _, personalWorkspaces, _, err := base.WorkspaceContent(l.ctx, community.FolderUID)
		if err != nil {
			return nil, err
		}
		activeUIDs := make([]string, 0)
		for _, workspace := range personalWorkspaces {
			activeUIDs = append(activeUIDs, workspace.WorkspaceUID)
		}
		resp = &types.GetAuditPublicCommunityResponse{
			CommunityUID:         community.UID,
			CommunityTitle:       community.Title,
			CommunityDescription: community.Description,
			Status:               community.Status,
			FolderUID:            folder.UID,
			FolderName:           folder.FolderName,
			FolderNumber:         folder.FolderNumber,
			AuditCount:           count,
		}
		if err = copier.Copy(&resp.PersonalWorkspaces, &personalWorkspaces); err != nil {
			return nil, err
		}
	}
	return resp, nil
}
