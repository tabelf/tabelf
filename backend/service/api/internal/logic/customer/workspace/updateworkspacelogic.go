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

type UpdateWorkspaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateWorkspaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateWorkspaceLogic {
	return &UpdateWorkspaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateWorkspaceLogic) UpdateWorkspace(req *types.UpdateWorkspaceRequest) (resp *types.UpdateWorkspaceResponse, err error) {
	if app.IsBlank(req.WorkspaceName) {
		return nil, app.ErrCustomerWorkspaceNameEmpty(l.ctx)
	}
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	if err = app.EntClient.Workspace.Update().
		SetName(req.WorkspaceName).
		Where(
			entworkspace.UID(req.WorkspaceUID),
		).Exec(l.ctx); err != nil {
		return nil, err
	}
	return &types.UpdateWorkspaceResponse{
		Message: app.HttpOK,
	}, nil
}
