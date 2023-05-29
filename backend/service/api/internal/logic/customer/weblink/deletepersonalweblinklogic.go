package weblink

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	entweblink "tabelf/backend/gen/entschema/weblink"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type DeletePersonalWebLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePersonalWebLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePersonalWebLinkLogic {
	return &DeletePersonalWebLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePersonalWebLinkLogic) DeletePersonalWebLink(req *types.DeletePersonalWebLinkRequest) (resp *types.DeletePersonalWebLinkResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	if err = app.EntClient.WebLink.Update().
		SetDeactivatedAt(time.Now()).
		Where(
			entweblink.UID(req.LinkUID),
			entweblink.WorkspaceUID(req.WorkspaceUID),
		).Exec(l.ctx); err != nil {
		return nil, err
	}
	return &types.DeletePersonalWebLinkResponse{
		Message: app.HttpOK,
	}, nil
}
