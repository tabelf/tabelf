package community

import (
	"context"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	entcommunity "tabelf/backend/gen/entschema/community"
	entcommunitymeta "tabelf/backend/gen/entschema/communitymeta"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSelfPublicCommunityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSelfPublicCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSelfPublicCommunityLogic {
	return &GetSelfPublicCommunityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSelfPublicCommunityLogic) GetSelfPublicCommunity(req *types.GetSelfPublicCommunityRequest) (resp *types.GetSelfPublicCommunityResponse, err error) {
	communities := make([]*entschema.Community, 0)
	switch req.Category {
	case app.SelfPublicCategory:
		communities, err = app.EntClient.Community.Query().Where(
			entcommunity.UserUID(req.UserUID),
			entcommunity.DeactivatedAtIsNil(),
		).All(l.ctx)
		if err != nil {
			return nil, err
		}
	case app.SelfStarCategory, app.SelfRecentCategory:
		metaQuery := app.EntClient.CommunityMeta.Query().Where(
			entcommunitymeta.UserUID(req.UserUID),
			entcommunitymeta.DeactivatedAtIsNil(),
		)
		if req.Category == app.SelfStarCategory {
			metaQuery.Where(entcommunitymeta.HasStar(true))
		} else {
			metaQuery.Where(entcommunitymeta.HasUsed(true))
		}
		metas, err := metaQuery.All(l.ctx)
		if err != nil {
			return nil, err
		}
		if len(metas) != 0 {
			communityUIDs := make([]string, 0)
			for _, meta := range metas {
				communityUIDs = append(communityUIDs, meta.CommunityUID)
				communities, err = app.EntClient.Community.Query().Where(
					entcommunity.UIDIn(communityUIDs...),
					entcommunity.DeactivatedAtIsNil(),
				).All(l.ctx)
				if err != nil {
					return nil, err
				}
			}
		}
	case app.SelfAuditCategory:
	default:
		return nil, app.ErrCommunityCategoryInvalid(l.ctx)
	}
	hasAdmin, err := base.HasAdminAuthority(l.ctx, req.UserUID)
	if err != nil {
		return nil, err
	}
	data := make([]*types.PublicCommunity, 0)
	if len(communities) == 0 {
		return &types.GetSelfPublicCommunityResponse{
			HasAdmin: hasAdmin,
			Data:     data,
		}, nil
	}
	userUIDs := make([]string, 0)
	for _, community := range communities {
		userUIDs = append(userUIDs, community.UserUID)
	}
	accounts, err := app.EntClient.Account.Query().Where(
		entaccount.UIDIn(userUIDs...),
		entaccount.DeactivatedAtIsNil(),
	).All(l.ctx)
	if err != nil {
		return nil, err
	}
	userMap := make(map[string]*entschema.Account)
	for _, account := range accounts {
		userMap[account.UID] = account
	}
	for _, community := range communities {
		userUID, userImage, userName := "", "", ""
		if user, ok := userMap[community.UserUID]; ok {
			userUID = user.UID
			userImage = user.Image
			userName = user.Nickname
		}
		data = append(data, &types.PublicCommunity{
			UID:         community.UID,
			Title:       community.Title,
			Description: community.Description,
			Image:       community.Image,
			Praise:      community.Praise,
			Used:        community.Used,
			View:        community.View,
			UserUID:     userUID,
			UserImage:   userImage,
			UserName:    userName,
			Status:      community.Status,
		})
	}
	return &types.GetSelfPublicCommunityResponse{
		HasAdmin: hasAdmin,
		Data:     data,
	}, nil
}
