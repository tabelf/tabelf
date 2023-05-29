package folder

import (
	"context"
	"golang.org/x/sync/errgroup"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	entcollaboration "tabelf/backend/gen/entschema/collaboration"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"

	"github.com/zeromicro/go-zero/core/logx"
)

type CopySharePersonalFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCopySharePersonalFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CopySharePersonalFolderLogic {
	return &CopySharePersonalFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CopySharePersonalFolderLogic) CopySharePersonalFolder(req *types.CopySharePersonalFolderRequest) (resp *types.CopySharePersonalFolderResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	collaboration, err := app.EntClient.Collaboration.Query().Where(
		entcollaboration.ShardUID(req.ShareUID),
		entcollaboration.UserUID(req.UserUID),
		entcollaboration.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		if entschema.IsNotFound(err) {
			return nil, app.ErrPersonalFolderCollNotExist(l.ctx)
		}
		return nil, err
	}
	var (
		g       errgroup.Group
		folder  *entschema.PersonalFolder
		account *entschema.Account
	)
	g.Go(func() (err error) {
		folder, err = app.EntClient.PersonalFolder.Query().Where(
			entpersonalfolder.UID(collaboration.FolderUID),
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
	return &types.CopySharePersonalFolderResponse{
		Message: app.HttpOK,
	}, nil
}
