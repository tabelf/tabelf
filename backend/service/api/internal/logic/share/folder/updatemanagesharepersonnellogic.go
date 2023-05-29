package folder

import (
	"context"
	entcollaboration "tabelf/backend/gen/entschema/collaboration"
	entsharelink "tabelf/backend/gen/entschema/sharelink"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateManageSharePersonnelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateManageSharePersonnelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateManageSharePersonnelLogic {
	return &UpdateManageSharePersonnelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateManageSharePersonnelLogic) UpdateManageSharePersonnel(req *types.UpdateManageSharePersonnelRequest) (resp *types.UpdateManageSharePersonnelResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	if app.IsBlank(req.OtherUserUID) {
		return nil, app.ErrCustomerUIDEmpty(l.ctx)
	}
	if !app.StringSliceContains([]string{
		app.ShareReadAuthority,
		app.ShareEditAuthority,
	}, req.Authority) {
		return nil, app.ErrCustomerFolderShareAuthorityInvalid(l.ctx)
	}
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
		SetAuthority(req.Authority).
		Where(entcollaboration.ID(collaboration.ID)).
		Exec(l.ctx); err != nil {
		return nil, err
	}
	return &types.UpdateManageSharePersonnelResponse{
		Message: app.HttpOK,
	}, nil
}
