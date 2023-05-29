package weblink

import (
	"context"
	"tabelf/backend/gen/entschema"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"
	entweblink "tabelf/backend/gen/entschema/weblink"
	entworkspace "tabelf/backend/gen/entschema/workspace"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RestoreDeletePersonalWebLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRestoreDeletePersonalWebLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RestoreDeletePersonalWebLinkLogic {
	return &RestoreDeletePersonalWebLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RestoreDeletePersonalWebLinkLogic) RestoreDeletePersonalWebLink(req *types.RestoreDeletePersonalWebLinkRequest) (resp *types.RestoreDeletePersonalWebLinkResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	delLink, err := app.EntClient.WebLink.Query().Where(
		entweblink.UID(req.LinkUID),
		entweblink.DeactivatedAtNotNil(),
		entweblink.ForeverDelete(false),
	).First(l.ctx)
	if err != nil {
		return nil, err
	}
	if err = app.WithTx(l.ctx, app.EntClient, func(tx *entschema.Tx) (err error) {
		if err = app.EntClient.WebLink.Update().
			ClearDeactivatedAt().
			Where(entweblink.ID(delLink.ID)).
			Exec(l.ctx); err != nil {
			return err
		}
		workspace, err := app.EntClient.Workspace.Query().Where(
			entworkspace.UID(delLink.WorkspaceUID),
		).First(l.ctx)
		if err != nil {
			return err
		}
		if workspace.DeactivatedAt != nil {
			if err = app.EntClient.Workspace.Update().
				ClearDeactivatedAt().
				Where(entworkspace.UID(delLink.WorkspaceUID)).
				Exec(l.ctx); err != nil {
				return err
			}
		}
		folder, err := app.EntClient.PersonalFolder.Query().Where(
			entpersonalfolder.UID(workspace.PersonalFolderUID),
		).First(l.ctx)
		if err != nil {
			return err
		}
		if folder.DeactivatedAt != nil {
			if err = app.EntClient.PersonalFolder.Update().
				ClearDeactivatedAt().
				Where(entpersonalfolder.UID(workspace.PersonalFolderUID)).
				Exec(l.ctx); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &types.RestoreDeletePersonalWebLinkResponse{
		Message: app.HttpOK,
	}, nil
}
