package folder

import (
	"context"
	"tabelf/backend/gen/entschema"
	entcollaboration "tabelf/backend/gen/entschema/collaboration"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSharePersonalFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSharePersonalFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSharePersonalFolderLogic {
	return &UpdateSharePersonalFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSharePersonalFolderLogic) UpdateSharePersonalFolder(req *types.UpdateSharePersonalFolderRequest) (resp *types.UpdateSharePersonalFolderResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	if app.IsBlank(req.FolderName) {
		return nil, app.ErrPersonalFolderNameEmpty(l.ctx)
	}
	collaboration, err := app.EntClient.Collaboration.Query().Where(
		entcollaboration.ShardUID(req.ShareUID),
		entcollaboration.UserUID(req.UserUID),
		entcollaboration.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		if entschema.IsNotFound(err) {
			return nil, app.ErrAccountOutOfAuthority(l.ctx)
		}
		return nil, err
	}
	if collaboration.Authority == app.ShareReadAuthority {
		return nil, app.ErrCustomerShareOutOfEditAuthority(l.ctx)
	}
	if err = app.EntClient.PersonalFolder.Update().
		SetFolderName(req.FolderName).
		Where(entpersonalfolder.UID(collaboration.FolderUID)).
		Exec(l.ctx); err != nil {
		return nil, err
	}
	return &types.UpdateSharePersonalFolderResponse{
		Message: app.HttpOK,
	}, nil
}
