package community

import (
	"context"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	entcommunity "tabelf/backend/gen/entschema/community"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublicCommunityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPublicCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublicCommunityLogic {
	return &GetPublicCommunityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPublicCommunityLogic) GetPublicCommunity(req *types.GetPublicCommunityRequest) (resp *types.GetPublicCommunityResponse, err error) {
	communityQuery := app.EntClient.Community.Query().Where(
		entcommunity.Status(app.PassStatus),
		entcommunity.DeactivatedAtIsNil(),
	)
	if app.IsNotBlank(req.CategoryUID) {
		communityQuery.Where(entcommunity.CategoryUID(req.CategoryUID))
	}
	switch req.Sorted {
	case app.CommunityView: // 全部: 查看
		communityQuery.Order(entschema.Desc(entcommunity.FieldView))
	case app.CommunityPraise: // 点赞最多
		communityQuery.Order(entschema.Desc(entcommunity.FieldPraise))
	case app.CommunityUsed: // 使用最多
		communityQuery.Order(entschema.Desc(entcommunity.FieldUsed))
	case app.CommunityNew: // 最新
		communityQuery.Order(entschema.Desc(entcommunity.FieldCreatedAt))
	default:
		return nil, app.ErrCommunitySortedInvalid(l.ctx)
	}
	count, err := communityQuery.Count(l.ctx)
	if err != nil {
		return nil, err
	}
	data := make([]*types.PublicCommunity, 0)
	if count != 0 {
		communities, err := communityQuery.Offset(req.Offset).Limit(req.Limit).All(l.ctx)
		if err != nil {
			return nil, err
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
			})
		}
	}
	return &types.GetPublicCommunityResponse{
		Total: count,
		Data:  data,
	}, nil
}
