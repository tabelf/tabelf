package folder

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/sync/errgroup"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type CopyPersonalFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCopyPersonalFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CopyPersonalFolderLogic {
	return &CopyPersonalFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CopyPersonalFolderLogic) CopyPersonalFolder(req *types.CopyPersonalFolderRequest) (resp *types.CopyPersonalFolderResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	var (
		g       errgroup.Group
		folder  *entschema.PersonalFolder
		account *entschema.Account
	)
	g.Go(func() (err error) {
		folder, err = app.EntClient.PersonalFolder.Query().Where(
			entpersonalfolder.UID(req.FolderUID),
			entpersonalfolder.UserUID(req.UserUID),
			entpersonalfolder.DeactivatedAtIsNil(),
		).First(l.ctx)
		if err != nil {
			if entschema.IsNotFound(err) {
				return app.ErrPersonalFolderNotExist(l.ctx)
			}
			return err
		}
		return nil
	})
	g.Go(func() (err error) {
		account, err = app.EntClient.Account.Query().Where(
			entaccount.UID(req.UserUID),
			entaccount.DeactivatedAtIsNil(),
		).First(l.ctx)
		if err != nil {
			if entschema.IsNotFound(err) {
				return app.ErrAccountInvalid(l.ctx)
			}
			return err
		}
		return nil
	})
	if err = g.Wait(); err != nil {
		return nil, err
	}
	if _, err = base.CopyFolderContent(l.ctx, account, folder); err != nil {
		return nil, err
	}
	return &types.CopyPersonalFolderResponse{
		Message: app.HttpOK,
	}, nil
}
