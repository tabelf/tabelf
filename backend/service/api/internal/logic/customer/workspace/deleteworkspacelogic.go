package workspace

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	entweblink "tabelf/backend/gen/entschema/weblink"
	entworkspace "tabelf/backend/gen/entschema/workspace"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type DeleteWorkspaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteWorkspaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteWorkspaceLogic {
	return &DeleteWorkspaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteWorkspaceLogic) DeleteWorkspace(req *types.DeleteWorkspaceRequest) (resp *types.DeleteWorkspaceResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	if err = app.EntClient.Workspace.Update().
		SetDeactivatedAt(time.Now()).
		Where(
			entworkspace.UID(req.WorkspaceUID),
			entworkspace.PersonalFolderUID(req.FolderUID),
		).Exec(l.ctx); err != nil {
		return nil, err
	}
	if err = app.EntClient.WebLink.Update().
		SetDeactivatedAt(time.Now()).
		Where(
			entweblink.WorkspaceUID(req.WorkspaceUID),
			entweblink.DeactivatedAtIsNil(),
		).Exec(l.ctx); err != nil {
		return nil, err
	}
	return &types.DeleteWorkspaceResponse{
		Message: app.HttpOK,
	}, nil
}
