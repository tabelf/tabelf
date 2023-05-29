package folder

import (
	"context"
	entcollaboration "tabelf/backend/gen/entschema/collaboration"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"
	"time"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExitSharePersonalFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExitSharePersonalFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExitSharePersonalFolderLogic {
	return &ExitSharePersonalFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExitSharePersonalFolderLogic) ExitSharePersonalFolder(req *types.ExitSharePersonalFolderRequest) (resp *types.ExitSharePersonalFolderResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	collaboration, err := app.EntClient.Collaboration.Query().Where(
		entcollaboration.ShardUID(req.ShareUID),
		entcollaboration.UserUID(req.UserUID),
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
	return &types.ExitSharePersonalFolderResponse{
		Message: app.HttpOK,
	}, nil
}
