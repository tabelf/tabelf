package weblink

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	entweblink "tabelf/backend/gen/entschema/weblink"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type UpdatePersonalWebLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePersonalWebLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePersonalWebLinkLogic {
	return &UpdatePersonalWebLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePersonalWebLinkLogic) UpdatePersonalWebLink(req *types.UpdatePersonalWebLinkRequest) (resp *types.UpdatePersonalWebLinkResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	if err = app.EntClient.WebLink.Update().
		SetTitle(req.Title).
		SetDescription(req.Description).
		Where(
			entweblink.UID(req.LinkUID),
			entweblink.WorkspaceUID(req.WorkspaceUID),
		).Exec(l.ctx); err != nil {
		return nil, err
	}
	return &types.UpdatePersonalWebLinkResponse{
		Message: app.HttpOK,
	}, nil
}
