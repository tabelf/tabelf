package folder

import (
	"context"
	"fmt"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	entcollaboration "tabelf/backend/gen/entschema/collaboration"
	entsharelink "tabelf/backend/gen/entschema/sharelink"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"
	"time"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SharePersonalFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSharePersonalFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SharePersonalFolderLogic {
	return &SharePersonalFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SharePersonalFolderLogic) SharePersonalFolder(req *types.SharePersonalFolderRequest) (resp *types.SharePersonalFolderResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	// 查询该文件的协作链接
	shareLink, err := app.EntClient.ShareLink.Query().Where(
		entsharelink.FolderUID(req.FolderUID),
		entsharelink.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		if entschema.IsNotFound(err) {
			return nil, app.ErrPersonalFolderShareLinkNotExist(l.ctx)
		}
		return nil, err
	}
	// 查询该文件的所有者
	owner, err := app.EntClient.Account.Query().Where(
		entaccount.UID(shareLink.UserUID),
		entaccount.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		if entschema.IsNotFound(err) {
			return nil, app.ErrAccountInvalid(l.ctx)
		}
		return nil, err
	}
	// 检查最近更新时间
	if app.GetYMD(shareLink.RecentAt) != app.GetYMD(time.Now()) {
		expiredAt := time.Now().AddDate(0, 0, shareLink.ValidDay)
		if shareLink.ValidDay == -1 {
			expiredAt = app.ParseTime(app.ForEverValid)
		}
		if err = app.EntClient.ShareLink.Update().
			SetExpiredAt(expiredAt).
			SetRecentAt(time.Now()).
			Exec(l.ctx); err != nil {
			return nil, err
		}
	}
	// 检查协作人数
	count, err := app.EntClient.Collaboration.Query().Where(
		entcollaboration.ShardUID(shareLink.UID),
		entcollaboration.DeactivatedAtIsNil(),
	).Count(l.ctx)
	if err != nil {
		return nil, err
	}
	sharePersonnels := make([]*types.SharePersonnel, 0)
	if req.Offset == 0 {
		sharePersonnels = append(sharePersonnels, &types.SharePersonnel{
			UserUID:       owner.UID,
			UserName:      owner.Nickname,
			Email:         owner.Email,
			Authority:     "1",
			HasMembership: base.MembershipValidity(owner),
			Image:         owner.Image,
			HasSelf:       req.UserUID == owner.UID,
			Sequence:      0,
		})
		count++
	}
	if count != 0 {
		collaborations, err := app.EntClient.Collaboration.Query().Where(
			entcollaboration.ShardUID(shareLink.UID),
			entcollaboration.DeactivatedAtIsNil(),
		).Offset(req.Offset).Limit(req.Limit).All(l.ctx)
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
		for i, account := range accounts {
			sharePersonnels = append(sharePersonnels, &types.SharePersonnel{
				CollUID:       userCollaborationMap[account.UID].UID,
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
	}
	return &types.SharePersonalFolderResponse{
		FolderUID:      shareLink.FolderUID,
		ShareUID:       shareLink.UID,
		ShareLink:      fmt.Sprintf("%s/v/%s", app.Basic.Domain, shareLink.UID),
		Authority:      shareLink.Authority,
		ExpiredDay:     shareLink.ValidDay,
		HasOwner:       req.UserUID == owner.UID,
		ShareNum:       count,
		SharePersonnel: sharePersonnels,
	}, nil
}
