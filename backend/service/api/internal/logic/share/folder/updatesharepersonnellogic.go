package folder

import (
	"context"
	entcollaboration "tabelf/backend/gen/entschema/collaboration"
	entsharelink "tabelf/backend/gen/entschema/sharelink"
	"tabelf/backend/service/app"
	"time"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSharePersonnelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSharePersonnelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSharePersonnelLogic {
	return &UpdateSharePersonnelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSharePersonnelLogic) UpdateSharePersonnel(req *types.UpdateSharePersonnelRequest) (resp *types.UpdateSharePersonnelResponse, err error) {
	link, err := app.EntClient.ShareLink.Query().Where(
		entsharelink.UID(req.ShareUID),
		entsharelink.UserUID(req.UserUID),
		entsharelink.FolderUID(req.FolderUID),
		entsharelink.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		return nil, err
	}
	collaboration, err := app.EntClient.Collaboration.Query().Where(
		entcollaboration.UID(req.CollUID),
		entcollaboration.ShardUID(link.UID),
		entcollaboration.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		return nil, err
	}
	switch req.Type {
	case app.ShareReadAuthority, app.ShareEditAuthority: // 可查看, 可编辑
		if err = app.EntClient.Collaboration.Update().
			SetAuthority(req.Type).
			Where(entcollaboration.UID(collaboration.UID)).
			Exec(l.ctx); err != nil {
			return nil, err
		}
	case app.ShareRemoveAuthority: // 移除
		if err = app.EntClient.Collaboration.Update().
			SetDeactivatedAt(time.Now()).
			Where(entcollaboration.UID(collaboration.UID)).
			Exec(l.ctx); err != nil {
			return nil, err
		}
	}
	return &types.UpdateSharePersonnelResponse{
		Message: app.HttpOK,
	}, nil
}
