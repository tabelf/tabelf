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

type DeleteSharePersonnelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSharePersonnelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSharePersonnelLogic {
	return &DeleteSharePersonnelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSharePersonnelLogic) DeleteSharePersonnel(req *types.DeleteSharePersonnelRequest) (resp *types.DeleteSharePersonnelResponse, err error) {
	shareLink, err := app.EntClient.ShareLink.Query().Where(
		entsharelink.UserUID(req.UserUID),
		entsharelink.UID(req.ShareUID),
		entsharelink.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		return nil, err
	}
	collaboration, err := app.EntClient.Collaboration.Query().Where(
		entcollaboration.ShardUID(req.ShareUID),
		entcollaboration.FolderUID(shareLink.FolderUID),
		entcollaboration.UserUID(req.OtherUserUID),
		entcollaboration.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		return nil, err
	}
	if err = app.EntClient.Collaboration.Update().
		SetDeactivatedAt(time.Now()).
		Where(entcollaboration.ID(collaboration.ID)).
		Exec(l.ctx); err != nil {
		return nil, err
	}
	return &types.DeleteSharePersonnelResponse{
		Message: app.HttpOK,
	}, nil
}
