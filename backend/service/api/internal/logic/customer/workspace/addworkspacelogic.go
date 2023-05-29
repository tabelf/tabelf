package workspace

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type AddWorkspaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddWorkspaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddWorkspaceLogic {
	return &AddWorkspaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddWorkspaceLogic) AddWorkspace(req *types.AddWorkspaceRequest) (resp *types.AddWorkspaceResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	if app.IsBlank(req.WorkspaceName) {
		return nil, app.ErrCustomerWorkspaceNameEmpty(l.ctx)
	}
	if err = app.EntClient.Workspace.Create().
		SetName(req.WorkspaceName).
		SetType(app.PersonalWorkspaceType).
		SetPersonalFolderUID(req.FolderUID).
		SetUserUID(req.UserUID).
		SetIsOpen(true).
		Exec(l.ctx); err != nil {
		return nil, err
	}
	return &types.AddWorkspaceResponse{
		Message: app.HttpOK,
	}, nil
}
