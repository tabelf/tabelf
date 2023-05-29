package community

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"

	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type CreatePublicCommunityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePublicCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePublicCommunityLogic {
	return &CreatePublicCommunityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePublicCommunityLogic) CreatePublicCommunity(req *types.CreatePublicCommunityRequest) (resp *types.CreatePublicCommunityResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	if err = l.ValidatePublic(req); err != nil {
		return nil, err
	}
	account, err := app.EntClient.Account.Query().Where(
		entaccount.UID(req.UserUID),
		entaccount.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		if entschema.IsNotFound(err) {
			return nil, app.ErrAccountInvalid(l.ctx)
		}
		return nil, err
	}
	exist, err := app.EntClient.PersonalFolder.Query().Where(
		entpersonalfolder.UID(req.FolderUID),
		entpersonalfolder.UserUID(req.UserUID),
		entpersonalfolder.HasOpen(true),
		entpersonalfolder.DeactivatedAtIsNil(),
	).Exist(l.ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, app.ErrCommunityRepeatPublic(l.ctx)
	}
	imageName := fmt.Sprintf("%s-%s.png", account.UID, req.FolderUID)

	if err = app.WithTx(l.ctx, app.EntClient, func(tx *entschema.Tx) (err error) {
		if err = tx.Community.Create().
			SetTitle(req.Title).
			SetDescription(req.Description).
			SetImage(app.TxCosCommunityURL + "/" + imageName).
			SetTags(req.Tags).
			SetFolderUID(req.FolderUID).
			SetUserUID(req.UserUID).
			SetStatus(app.WaitAuditStatus).
			SetCategoryUID(req.CategoryUID).
			Exec(l.ctx); err != nil {
			return err
		}
		if err = tx.PersonalFolder.Update().SetHasOpen(true).
			Where(
				entpersonalfolder.UID(req.FolderUID),
				entpersonalfolder.UserUID(req.UserUID),
				entpersonalfolder.DeactivatedAtIsNil(),
			).Exec(l.ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &types.CreatePublicCommunityResponse{
		Message: app.HttpOK,
	}, nil
}

func (l *CreatePublicCommunityLogic) ValidatePublic(req *types.CreatePublicCommunityRequest) error {
	if app.IsBlank(req.CategoryUID) {
		return app.ErrCommunityCategoryEmpty(l.ctx)
	}
	if app.IsBlank(req.Title) {
		return app.ErrCommunityTitleEmpty(l.ctx)
	}
	if app.IsBlank(req.Description) {
		return app.ErrCommunityDescriptionEmpty(l.ctx)
	}
	if app.IsBlank(req.FolderUID) {
		return app.ErrCommunityFolderUIDEmpty(l.ctx)
	}
	return nil
}
