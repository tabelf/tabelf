package community

import (
	"context"
	"golang.org/x/sync/errgroup"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	entcommunity "tabelf/backend/gen/entschema/community"
	entcommunitymeta "tabelf/backend/gen/entschema/communitymeta"
	entfocus "tabelf/backend/gen/entschema/focus"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublicCommunityDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPublicCommunityDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublicCommunityDetailLogic {
	return &GetPublicCommunityDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPublicCommunityDetailLogic) GetPublicCommunityDetail(req *types.GetPublicCommunityDetailRequest) (resp *types.GetPublicCommunityDetailResponse, err error) {
	var (
		g         errgroup.Group
		user      *entschema.Account
		hasPraise bool
		hasStar   bool
		openCount int
		hasFollow bool
	)
	community, err := app.EntClient.Community.Query().Where(
		entcommunity.UID(req.CommunityUID),
		entcommunity.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		return nil, err
	}
	if app.IsNotBlank(req.UserUID) {
		g.Go(func() (err error) {
			// 查询我自己的对该社区对元数据
			meta, err := app.EntClient.CommunityMeta.Query().Where(
				entcommunitymeta.UserUID(req.UserUID),
				entcommunitymeta.CommunityUID(req.CommunityUID),
				entcommunitymeta.DeactivatedAtIsNil(),
			).First(l.ctx)
			if err != nil && !entschema.IsNotFound(err) {
				return err
			}
			if meta != nil {
				hasPraise = meta.HasPraise
				hasStar = meta.HasStar
			}
			return nil
		})
		g.Go(func() (err error) {
			// 查询我自己的对该社区用户的关注数据
			focus, err := app.EntClient.Focus.Query().Where(
				entfocus.FollowerUID(req.UserUID),
				entfocus.FolloweeUID(community.UserUID),
				entfocus.DeactivatedAtIsNil(),
			).First(l.ctx)
			if err != nil && !entschema.IsNotFound(err) {
				return err
			}
			if focus == nil {
				hasFollow = false
			} else {
				hasFollow = focus.Status
			}
			return nil
		})
	}
	g.Go(func() (err error) {
		// 查询社区的用户
		user, err = app.EntClient.Account.Query().Where(
			entaccount.UID(community.UserUID),
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
	g.Go(func() (err error) {
		// 查询社区用户的发布数量
		openCount, err = app.EntClient.PersonalFolder.Query().Where(
			entpersonalfolder.UserUID(community.UserUID),
			entpersonalfolder.DeactivatedAtIsNil(),
			entpersonalfolder.HasOpen(true),
		).Count(l.ctx)
		return err
	})
	if err = g.Wait(); err != nil {
		return nil, err
	}
	resp = &types.GetPublicCommunityDetailResponse{
		UID:         community.UID,
		Title:       community.Title,
		Description: community.Description,
		HtmlDesc:    app.TruncateString(app.HtmlPlainText(community.Description), 180),
		Image:       community.Image,
		Tags:        community.Tags,
		Praise:      community.Praise,
		HasPraise:   hasPraise,
		Used:        community.Used,
		View:        community.View,
		Star:        community.Star,
		HasStar:     hasStar,
		UserUID:     user.UID,
		UserImage:   user.Image,
		UserName:    user.Nickname,
		Open:        openCount,
		Fans:        user.Fans,
		HasSelf:     req.UserUID == community.UserUID,
		HasFollow:   hasFollow,
	}
	return resp, nil
}
