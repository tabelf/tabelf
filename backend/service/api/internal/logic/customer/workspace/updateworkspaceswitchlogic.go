package workspace

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	entworkspace "tabelf/backend/gen/entschema/workspace"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type UpdateWorkspaceSwitchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateWorkspaceSwitchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateWorkspaceSwitchLogic {
	return &UpdateWorkspaceSwitchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateWorkspaceSwitchLogic) UpdateWorkspaceSwitch(req *types.UpdateWorkspaceSwitchRequest) (resp *types.UpdateWorkspaceSwitchResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	if len(req.ActiveWorkspaceUIDs) == 0 {
		if err = app.EntClient.Workspace.Update().
			SetIsOpen(false).
			Where(
				entworkspace.Type(app.PersonalWorkspaceType),
				entworkspace.PersonalFolderUID(req.FolderUID),
				entworkspace.DeactivatedAtIsNil(),
			).Exec(l.ctx); err != nil {
			return nil, err
		}
	} else {
		if err = app.EntClient.Workspace.Update().
			SetIsOpen(false).
			Where(
				entworkspace.Type(app.PersonalWorkspaceType),
				entworkspace.PersonalFolderUID(req.FolderUID),
				entworkspace.UIDNotIn(req.ActiveWorkspaceUIDs...),
				entworkspace.DeactivatedAtIsNil(),
			).Exec(l.ctx); err != nil {
			return nil, err
		}
		if err = app.EntClient.Workspace.Update().
			SetIsOpen(true).
			Where(
				entworkspace.Type(app.PersonalWorkspaceType),
				entworkspace.PersonalFolderUID(req.FolderUID),
				entworkspace.UIDIn(req.ActiveWorkspaceUIDs...),
				entworkspace.DeactivatedAtIsNil(),
			).Exec(l.ctx); err != nil {
			return nil, err
		}
	}
	return &types.UpdateWorkspaceSwitchResponse{
		Message: app.HttpOK,
	}, nil
}
