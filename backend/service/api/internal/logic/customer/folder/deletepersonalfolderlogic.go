package folder

import (
	"context"
	"tabelf/backend/gen/entschema"
	entweblink "tabelf/backend/gen/entschema/weblink"
	entworkspace "tabelf/backend/gen/entschema/workspace"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type DeletePersonalFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePersonalFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePersonalFolderLogic {
	return &DeletePersonalFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePersonalFolderLogic) DeletePersonalFolder(req *types.DeletePersonalFolderRequest) (resp *types.DeletePersonalFolderResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	workspaces, err := app.EntClient.Workspace.Query().Where(
		entworkspace.UserUID(req.UserUID),
		entworkspace.PersonalFolderUID(req.FolderUID),
		entworkspace.DeactivatedAtIsNil(),
	).All(l.ctx)
	if err != nil {
		return nil, err
	}
	workspaceUIDs := make([]string, len(workspaces))
	for i, workspace := range workspaces {
		workspaceUIDs[i] = workspace.UID
	}
	if err = app.WithTx(l.ctx, app.EntClient, func(tx *entschema.Tx) (err error) {
		if err = tx.PersonalFolder.Update().
			SetDeactivatedAt(time.Now()).
			Where(
				entpersonalfolder.UID(req.FolderUID),
				entpersonalfolder.UserUID(req.UserUID),
			).Exec(l.ctx); err != nil {
			return err
		}
		if len(workspaceUIDs) != 0 {
			if err = tx.Workspace.Update().
				SetDeactivatedAt(time.Now()).
				Where(
					entworkspace.UIDIn(workspaceUIDs...),
				).Exec(l.ctx); err != nil {
				return err
			}
			if err = tx.WebLink.Update().
				SetDeactivatedAt(time.Now()).
				Where(
					entweblink.WorkspaceUIDIn(workspaceUIDs...),
					entweblink.DeactivatedAtIsNil(),
				).Exec(l.ctx); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &types.DeletePersonalFolderResponse{
		Message: app.HttpOK,
	}, nil
}
