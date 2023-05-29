package folder

import (
	"context"
	"tabelf/backend/gen/entschema"
	entcollaboration "tabelf/backend/gen/entschema/collaboration"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"
	entsharelink "tabelf/backend/gen/entschema/sharelink"
	entweblink "tabelf/backend/gen/entschema/weblink"
	entworkspace "tabelf/backend/gen/entschema/workspace"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"
	"time"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSharePersonalFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSharePersonalFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSharePersonalFolderLogic {
	return &DeleteSharePersonalFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSharePersonalFolderLogic) DeleteSharePersonalFolder(req *types.DeleteSharePersonalFolderRequest) (resp *types.DeleteSharePersonalFolderResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	collaboration, err := app.EntClient.Collaboration.Query().Where(
		entcollaboration.ShardUID(req.ShareUID),
		entcollaboration.UserUID(req.UserUID),
		entcollaboration.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		if entschema.IsNotFound(err) {
			return nil, app.ErrAccountOutOfAuthority(l.ctx)
		}
		return nil, err
	}
	if collaboration.Authority == app.ShareReadAuthority {
		return nil, app.ErrCustomerShareOutOfEditAuthority(l.ctx)
	}
	workspaces, err := app.EntClient.Workspace.Query().Where(
		entworkspace.PersonalFolderUID(collaboration.FolderUID),
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
			Where(entpersonalfolder.UID(collaboration.FolderUID)).
			Exec(l.ctx); err != nil {
			return err
		}
		if len(workspaceUIDs) != 0 {
			if err = tx.Workspace.Update().
				SetDeactivatedAt(time.Now()).
				Where(entworkspace.UIDIn(workspaceUIDs...)).
				Exec(l.ctx); err != nil {
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
		// 删除分享链接
		if err = tx.ShareLink.Update().
			SetDeactivatedAt(time.Now()).
			Where(
				entsharelink.FolderUID(collaboration.FolderUID),
				entsharelink.DeactivatedAtIsNil(),
			).Exec(l.ctx); err != nil {
			return err
		}
		// 删除加入的协作人
		if err = tx.Collaboration.Update().
			SetDeactivatedAt(time.Now()).
			Where(
				entcollaboration.FolderUID(collaboration.FolderUID),
				entcollaboration.DeactivatedAtIsNil(),
			).Exec(l.ctx); err != nil {
			return nil
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &types.DeleteSharePersonalFolderResponse{
		Message: app.HttpOK,
	}, nil
}
