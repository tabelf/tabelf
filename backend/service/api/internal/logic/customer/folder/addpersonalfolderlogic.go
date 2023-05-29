package folder

import (
	"context"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"
	"tabelf/backend/spec/schema"
	"time"
	"unicode/utf8"

	"github.com/zeromicro/go-zero/core/logx"

	"tabelf/backend/gen/entschema"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type AddPersonalFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPersonalFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPersonalFolderLogic {
	return &AddPersonalFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPersonalFolderLogic) AddPersonalFolder(req *types.AddPersonalFolderRequest) (resp *types.AddPersonalFolderResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	if app.IsBlank(req.FolderName) {
		return nil, app.ErrPersonalFolderNameEmpty(l.ctx)
	}
	if utf8.RuneCountInString(req.FolderName) > schema.LenFolderName {
		return nil, app.ErrPersonalFolderNameLimit(l.ctx)
	}
	folderNumber := ""
	for {
		folderNumber = app.RandomString(app.FolderNumberCount)
		exist, err := app.EntClient.PersonalFolder.Query().Where(
			entpersonalfolder.FolderNumber(folderNumber),
			entpersonalfolder.UserUID(req.UserUID),
			entpersonalfolder.DeactivatedAtIsNil(),
		).Exist(l.ctx)
		if err != nil {
			return nil, err
		}
		if !exist {
			break
		}
	}
	var folder *entschema.PersonalFolder
	// 创建个人文件，并创建该文件对应的分享链接
	if err = app.WithTx(l.ctx, app.EntClient, func(tx *entschema.Tx) error {
		if folder, err = tx.PersonalFolder.Create().
			SetFolderName(req.FolderName).
			SetUserUID(req.UserUID).
			SetFolderNumber(folderNumber).
			Save(l.ctx); err != nil {
			return err
		}
		if err = tx.Workspace.Create().
			SetName(app.DefaultWorkspaceName).
			SetType(app.PersonalWorkspaceType).
			SetPersonalFolderUID(folder.UID).
			SetUserUID(req.UserUID).
			SetIsOpen(true).
			Exec(l.ctx); err != nil {
			return err
		}
		return tx.ShareLink.Create().
			SetFolderUID(folder.UID).
			SetUserUID(req.UserUID).
			SetAuthority(app.ShareReadAuthority).
			SetValidDay(app.ForEverDay).
			SetExpiredAt(app.ParseTime(app.ForEverValid)).
			SetRecentAt(time.Now()).
			SetFolderNumber(folder.FolderNumber).
			Exec(l.ctx)
	}); err != nil {
		return nil, err
	}
	return &types.AddPersonalFolderResponse{
		FolderUID:    folder.UID,
		FolderNumber: folder.FolderNumber,
	}, nil
}
