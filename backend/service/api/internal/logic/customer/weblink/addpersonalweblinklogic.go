package weblink

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type AddPersonalWebLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPersonalWebLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPersonalWebLinkLogic {
	return &AddPersonalWebLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPersonalWebLinkLogic) AddPersonalWebLink(req *types.AddPersonalWebLinkRequest) (resp *types.AddPersonalWebLinkResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	if app.IsBlank(req.URL) {
		return nil, app.ErrCustomerWebLinkEmpty(l.ctx)
	}
	if err = base.AddWebLink(l.ctx, req.URL, req.Title, req.UserUID, req.WorkspaceUID); err != nil {
		return nil, err
	}
	return &types.AddPersonalWebLinkResponse{
		Message: app.HttpOK,
	}, nil
}
