package community

import (
	"context"
	"golang.org/x/sync/errgroup"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	entcommunity "tabelf/backend/gen/entschema/community"
	entcommunitymeta "tabelf/backend/gen/entschema/communitymeta"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePublicCommunityMetaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePublicCommunityMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePublicCommunityMetaLogic {
	return &UpdatePublicCommunityMetaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePublicCommunityMetaLogic) UpdatePublicCommunityMeta(req *types.UpdatePublicCommunityMetaRequest) (resp *types.UpdatePublicCommunityMetaResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	community, err := app.EntClient.Community.Query().Where(
		entcommunity.UID(req.CommunityUID),
		entcommunity.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		return nil, err
	}
	meta, err := app.EntClient.CommunityMeta.Query().Where(
		entcommunitymeta.UserUID(req.UserUID),
		entcommunitymeta.CommunityUID(req.CommunityUID),
		entcommunitymeta.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil && !entschema.IsNotFound(err) {
		return nil, err
	}
	hasUsed := false
	folderNumber := ""
	if err = app.WithTx(l.ctx, app.EntClient, func(tx *entschema.Tx) (err error) {
		communityUpdate := tx.Community.Update()
		if meta == nil {
			metaCreate := tx.CommunityMeta.Create().
				SetUserUID(req.UserUID).
				SetCommunityUID(req.CommunityUID)
			switch req.MetaType {
			case app.CommunityView: // 查看
				metaCreate.SetHasView(true)
				communityUpdate.AddView(1)
			case app.CommunityPraise: // 点赞
				metaCreate.SetHasPraise(true)
				communityUpdate.AddPraise(1)
			case app.CommunityStar: // 收藏
				metaCreate.SetHasStar(true)
				communityUpdate.AddStar(1)
			case app.CommunityUsed:
				metaCreate.SetHasUsed(true)
				communityUpdate.AddUsed(1)
				hasUsed = true
			default:
				return app.ErrCommunitySortedInvalid(l.ctx)
			}
			if hasUsed {
				folderNumber, err = l.CopyPublicCommunity(req.UserUID, community.FolderUID)
				if err != nil {
					return err
				}
			}
			if err = metaCreate.Exec(l.ctx); err != nil {
				return err
			}
		} else {
			metaUpdate := tx.CommunityMeta.Update()
			switch req.MetaType {
			case app.StationView: // 查看
				if !meta.HasView {
					metaUpdate.SetHasView(true)
					communityUpdate.AddView(1)
				}
			case app.CommunityPraise: // 点赞
				if meta.HasPraise {
					metaUpdate.SetHasPraise(false)
					communityUpdate.AddPraise(-1)
				} else {
					metaUpdate.SetHasPraise(true)
					communityUpdate.AddPraise(1)
				}
			case app.CommunityStar: // 收藏
				if meta.HasStar {
					metaUpdate.SetHasStar(false)
					communityUpdate.AddStar(-1)
				} else {
					metaUpdate.SetHasStar(true)
					communityUpdate.AddStar(1)
				}
			case app.CommunityUsed: // 使用
				hasUsed = true
				if !meta.HasUsed {
					metaUpdate.SetHasUsed(true)
					communityUpdate.AddUsed(1)
				}
			default:
				return app.ErrCommunitySortedInvalid(l.ctx)
			}
			if hasUsed {
				folderNumber, err = l.CopyPublicCommunity(req.UserUID, community.FolderUID)
				if err != nil {
					return err
				}
			}
			if err = metaUpdate.
				Where(entcommunitymeta.ID(meta.ID)).
				Exec(l.ctx); err != nil {
				return err
			}
		}
		if err = communityUpdate.
			Where(entcommunity.ID(community.ID)).
			Exec(l.ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	community, err = app.EntClient.Community.Query().Where(
		entcommunity.ID(community.ID),
	).First(l.ctx)
	if err != nil {
		return nil, err
	}
	meta, err = app.EntClient.CommunityMeta.Query().Where(
		entcommunitymeta.UserUID(req.UserUID),
		entcommunitymeta.CommunityUID(req.CommunityUID),
		entcommunitymeta.DeactivatedAtIsNil(),
	).First(l.ctx)
	return &types.UpdatePublicCommunityMetaResponse{
		UID:          community.UID,
		FolderNumber: folderNumber,
		Praise:       community.Praise,
		HasPraise:    meta.HasPraise,
		Star:         community.Star,
		HasStar:      meta.HasStar,
		View:         community.View,
		Used:         community.Used,
		HasUsed:      meta.HasUsed,
	}, nil
}

func (l *UpdatePublicCommunityMetaLogic) CopyPublicCommunity(userUID string, folderUID string) (folderNumber string, err error) {
	var (
		g       errgroup.Group
		folder  *entschema.PersonalFolder
		account *entschema.Account
	)
	g.Go(func() (err error) {
		folder, err = app.EntClient.PersonalFolder.Query().Where(
			entpersonalfolder.UID(folderUID),
			entpersonalfolder.HasOpen(true),
			entpersonalfolder.DeactivatedAtIsNil(),
		).First(l.ctx)
		if err != nil {
			if entschema.IsNotFound(err) {
				return app.ErrPersonalFolderNotExist(l.ctx)
			}
			return err
		}
		return nil
	})
	g.Go(func() (err error) {
		account, err = app.EntClient.Account.Query().Where(
			entaccount.UID(userUID),
			entaccount.DeactivatedAtIsNil(),
		).First(l.ctx)
		if err != nil {
			if entschema.IsNotFound(err) {
				return app.ErrAccountInvalid(l.ctx)
			}
			return err
		}
		return nil
	})
	if err = g.Wait(); err != nil {
		return "", err
	}
	return base.CopyFolderContent(l.ctx, account, folder)
}
