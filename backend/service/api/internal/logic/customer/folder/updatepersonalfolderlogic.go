package folder

import (
	"context"
	"tabelf/backend/gen/entschema"
	"tabelf/backend/spec/schema"
	"unicode/utf8"

	"github.com/zeromicro/go-zero/core/logx"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type UpdatePersonalFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePersonalFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePersonalFolderLogic {
	return &UpdatePersonalFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePersonalFolderLogic) UpdatePersonalFolder(req *types.UpdatePersonalFolderRequest) (resp *types.UpdatePersonalFolderResponse, err error) {
	if app.IsBlank(req.FolderName) {
		return nil, app.ErrPersonalFolderNameEmpty(l.ctx)
	}
	if utf8.RuneCountInString(req.FolderName) > schema.LenFolderName {
		return nil, app.ErrPersonalFolderNameLimit(l.ctx)
	}
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	folder, err := app.EntClient.PersonalFolder.Query().Where(
		entpersonalfolder.UID(req.FolderUID),
		entpersonalfolder.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		if entschema.IsNotFound(err) {
			return nil, app.ErrPersonalFolderNotExist(l.ctx)
		}
		return nil, err
	}
	if err = app.EntClient.PersonalFolder.Update().
		SetFolderName(req.FolderName).
		Where(entpersonalfolder.ID(folder.ID)).
		Exec(l.ctx); err != nil {
		return nil, err
	}
	return &types.UpdatePersonalFolderResponse{
		Message: app.HttpOK,
	}, nil
}
