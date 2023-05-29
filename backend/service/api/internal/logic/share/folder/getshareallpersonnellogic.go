package folder

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	entcollaboration "tabelf/backend/gen/entschema/collaboration"
	entsharelink "tabelf/backend/gen/entschema/sharelink"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type GetShareAllPersonnelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShareAllPersonnelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShareAllPersonnelLogic {
	return &GetShareAllPersonnelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShareAllPersonnelLogic) GetShareAllPersonnel(req *types.GetShareAllPersonnelRequest) (resp *types.GetShareAllPersonnelResponse, err error) {
	shareLink, err := app.EntClient.ShareLink.Query().Where(
		entsharelink.UID(req.ShareUID),
		entsharelink.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		return nil, err
	}
	owner, err := app.EntClient.Account.Query().Where(
		entaccount.UID(shareLink.UserUID),
		entaccount.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		return nil, err
	}
	links, err := app.EntClient.ShareLink.Query().Where(
		entsharelink.FolderUID(shareLink.FolderUID),
		entsharelink.UserUID(shareLink.UserUID),
		entsharelink.DeactivatedAtIsNil(),
	).All(l.ctx)
	if err != nil {
		return nil, err
	}
	personnels := make([]*types.SharePersonnel, 0)
	if len(links) != 0 {
		linkUIDs := make([]string, len(links))
		for i, link := range links {
			linkUIDs[i] = link.UID
		}
		// 查询所有的协作者
		collaborations, err := app.EntClient.Collaboration.Query().Where(
			entcollaboration.ShardUIDIn(linkUIDs...),
			entcollaboration.DeactivatedAtIsNil(),
		).All(l.ctx)
		if err != nil {
			return nil, err
		}
		userUIDs := make([]string, 0)
		userCollaborationMap := make(map[string]*entschema.Collaboration, 0)
		for _, collaboration := range collaborations {
			userUIDs = append(userUIDs, collaboration.UserUID)
			userCollaborationMap[collaboration.UserUID] = collaboration
		}
		accounts, err := app.EntClient.Account.Query().Where(
			entaccount.UIDIn(userUIDs...),
			entaccount.DeactivatedAtIsNil(),
		).All(l.ctx)
		if err != nil {
			return nil, err
		}
		shares := make([]*types.SharePersonnel, 0)
		for i, account := range accounts {
			shares = append(shares, &types.SharePersonnel{
				UserUID:       account.UID,
				UserName:      account.Nickname,
				Email:         account.Email,
				Authority:     userCollaborationMap[account.UID].Authority,
				HasMembership: base.MembershipValidity(account),
				Image:         account.Image,
				HasSelf:       account.UID == req.UserUID,
				Sequence:      i + 1,
			})
		}
		personnels = append(personnels, &types.SharePersonnel{
			UserUID:   owner.UID,
			UserName:  owner.Nickname,
			Email:     owner.Email,
			Authority: "1",
			Image:     owner.Image,
			Sequence:  0,
		})
		personnels = append(personnels, shares...)
	}
	return &types.GetShareAllPersonnelResponse{
		HasOwner:       owner.UID == req.UserUID,
		SharePersonnel: personnels,
	}, nil
}
