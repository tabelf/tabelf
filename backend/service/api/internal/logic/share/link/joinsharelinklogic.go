package link

import (
	"context"
	"tabelf/backend/gen/entschema"
	entcollaboration "tabelf/backend/gen/entschema/collaboration"
	entsharelink "tabelf/backend/gen/entschema/sharelink"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"
	"time"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinShareLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJoinShareLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinShareLinkLogic {
	return &JoinShareLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JoinShareLinkLogic) JoinShareLink(req *types.JoinShareLinkRequest) (resp *types.JoinShareLinkResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	shareUID := req.ShareUID
	authority := ""
	if len(req.ShareUID) > 32 {
		authority = req.ShareUID[0:1]
		shareUID = req.ShareUID[1:]
	}
	shareLink, err := app.EntClient.ShareLink.Query().Where(
		entsharelink.UID(shareUID),
		entsharelink.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		if entschema.IsNotFound(err) {
			return nil, app.ErrCustomerFolderShareLinkNotExist(l.ctx)
		}
		return nil, err
	}
	if app.IsBlank(authority) { // 权限为空，表示通过复制链接进行邀请的，而非好友邀请。好友邀请不需要校验过期时间
		if time.Now().After(shareLink.ExpiredAt) {
			return nil, app.ErrCustomerFolderShareLinkExpired(l.ctx)
		}
	}
	if shareLink.UserUID == req.UserUID {
		return &types.JoinShareLinkResponse{
			HasShare:  false,
			FolderUID: shareLink.FolderUID,
		}, nil
	}
	if authority == "" {
		authority = shareLink.Authority
	}
	exist, err := app.EntClient.Collaboration.Query().Where(
		entcollaboration.UserUID(req.UserUID),
		entcollaboration.ShardUID(shareUID),
		entcollaboration.DeactivatedAtIsNil(),
	).Exist(l.ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return &types.JoinShareLinkResponse{
			HasShare:     true,
			FolderUID:    shareLink.FolderUID,
			FolderNumber: shareLink.FolderNumber,
		}, nil
	}
	if err = app.EntClient.Collaboration.Create().
		SetShardUID(shareUID).
		SetFolderUID(shareLink.FolderUID).
		SetUserUID(req.UserUID).
		SetAuthority(authority).
		SetFolderNumber(shareLink.FolderNumber).
		Exec(l.ctx); err != nil {
		return nil, err
	}
	return &types.JoinShareLinkResponse{
		HasShare:     true,
		FolderUID:    shareLink.FolderUID,
		FolderNumber: shareLink.FolderNumber,
	}, nil
}
