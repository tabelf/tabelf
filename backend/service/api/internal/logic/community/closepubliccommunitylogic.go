package community

import (
	"context"
	"tabelf/backend/gen/entschema"
	entcommunity "tabelf/backend/gen/entschema/community"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"
	"time"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClosePublicCommunityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClosePublicCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClosePublicCommunityLogic {
	return &ClosePublicCommunityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClosePublicCommunityLogic) ClosePublicCommunity(req *types.ClosePublicCommunityRequest) (resp *types.ClosePublicCommunityResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}

	community, err := app.EntClient.Community.Query().Where(
		entcommunity.UserUID(req.UserUID),
		entcommunity.UID(req.CommunityUID),
	).First(l.ctx)
	if err != nil {
		return nil, err
	}

	if community.DeactivatedAt != nil {
		return &types.ClosePublicCommunityResponse{
			Message: app.HttpOK,
		}, nil
	}

	if err = app.WithTx(l.ctx, app.EntClient, func(tx *entschema.Tx) (err error) {
		if err = tx.Community.Update().SetDeactivatedAt(time.Now()).
			Where(
				entcommunity.UserUID(req.UserUID),
				entcommunity.UID(req.CommunityUID),
				entcommunity.DeactivatedAtIsNil(),
			).Exec(l.ctx); err != nil {
			return err
		}
		if err = tx.PersonalFolder.Update().SetHasOpen(false).
			Where(
				entpersonalfolder.UID(community.FolderUID),
				entpersonalfolder.UserUID(req.UserUID),
				entpersonalfolder.DeactivatedAtIsNil(),
			).Exec(l.ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &types.ClosePublicCommunityResponse{
		Message: app.HttpOK,
	}, nil
}
