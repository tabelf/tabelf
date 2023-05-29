package station

import (
	"context"
	"golang.org/x/sync/errgroup"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"
	entstation "tabelf/backend/gen/entschema/station"
	entweblink "tabelf/backend/gen/entschema/weblink"
	entworkspace "tabelf/backend/gen/entschema/workspace"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodStationStarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodStationStarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodStationStarLogic {
	return &GoodStationStarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodStationStarLogic) GoodStationStar(req *types.GoodStationStarRequest) (resp *types.GoodStationStarResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	// 添加到该文件夹下到第一个工作空间，如果不存在给他创建一个
	workspace, err := GetFolderByFirstWorkspace(l.ctx, req.UserUID, req.FolderUID)
	if err != nil {
		return nil, err
	}
	var (
		g errgroup.Group
		//account *entschema.Account
		station *entschema.Station
		count   int
	)
	g.Go(func() (err error) {
		_, err = app.EntClient.Account.Query().Where(
			entaccount.UID(req.UserUID),
			entaccount.DeactivatedAtIsNil(),
		).First(l.ctx)
		if err != nil {
			if entschema.IsNotFound(err) {
				return app.ErrAccountInvalid(l.ctx)
			}
			return err
		}
		// 会员不做数量校验
		//if base.MembershipValidity(account) {
		//	return nil
		//}
		//if account.URLCount+1 > account.URLLimit {
		//	return app.ErrCustomerWebLinkLimit(l.ctx)
		//}
		return nil
	})
	g.Go(func() (err error) {
		station, err = app.EntClient.Station.Query().Where(
			entstation.UID(req.StationUID),
			entstation.DeactivatedAtIsNil(),
		).First(l.ctx)
		return err
	})
	g.Go(func() (err error) {
		count, err = app.EntClient.WebLink.Query().Where(
			entweblink.WorkspaceUID(workspace.UID),
			entweblink.DeactivatedAtIsNil(),
		).Count(l.ctx)
		return err
	})
	if err = g.Wait(); err != nil {
		return nil, err
	}
	if err = app.WithTx(l.ctx, app.EntClient, func(tx *entschema.Tx) (err error) {
		if err = tx.WebLink.Create().
			SetTitle(station.Title).
			SetDescription(station.Description).
			SetImage(station.Image).
			SetLink(station.Link).
			SetSequence(count).
			SetUserUID(req.UserUID).
			SetFileType(app.URLFileType).
			SetWorkspaceUID(workspace.UID).
			SetFolderUID(req.FolderUID).
			Exec(l.ctx); err != nil {
			return err
		}
		if err = tx.Account.Update().AddURLCount(1).
			Where(entaccount.UID(req.UserUID)).
			Exec(l.ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	meta, err := base.UpdateGoodStationMeta(l.ctx, req.StationUID, req.UserUID, app.StationStar)
	if err != nil {
		return nil, err
	}
	return &types.GoodStationStarResponse{
		UID:       meta.UID,
		Praise:    meta.Praise,
		HasPraise: meta.HasPraise,
		Star:      meta.Star,
		HasStar:   meta.HasStar,
		View:      meta.View,
	}, nil
}

func GetFolderByFirstWorkspace(ctx context.Context, userUID, folderUID string) (workspace *entschema.Workspace, err error) {
	folder, err := app.EntClient.PersonalFolder.Query().Where(
		entpersonalfolder.UID(folderUID),
		entpersonalfolder.UserUID(userUID),
		entpersonalfolder.DeactivatedAtIsNil(),
	).First(ctx)
	if err != nil {
		return nil, err
	}
	workspaces, err := app.EntClient.Workspace.Query().Where(
		entworkspace.PersonalFolderUID(folderUID),
		entworkspace.DeactivatedAtIsNil(),
	).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(workspaces) != 0 {
		return workspaces[0], nil
	}
	res, err := app.EntClient.Workspace.Create().
		SetName(app.DefaultWorkspaceName).
		SetType(app.PersonalWorkspaceType).
		SetPersonalFolderUID(folder.UID).
		SetUserUID(userUID).
		SetIsOpen(true).
		Save(ctx)
	return res, nil
}
