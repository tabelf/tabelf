package weblink

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	entweblink "tabelf/backend/gen/entschema/weblink"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type UpdateWorkspaceWebLinksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateWorkspaceWebLinksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateWorkspaceWebLinksLogic {
	return &UpdateWorkspaceWebLinksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateWorkspaceWebLinksLogic) UpdateWorkspaceWebLinks(req *types.UpdateWorkspaceWebLinksRequest) (resp *types.UpdateWorkspaceWebLinksResponse, err error) {
	for _, webLink := range req.WebLinks {
		if err = app.EntClient.WebLink.Update().
			SetSequence(webLink.Sequence).
			SetWorkspaceUID(req.WorkspaceUID).
			Where(
				entweblink.UID(webLink.LinkUID),
				entweblink.WorkspaceUID(req.OldWorkspaceUID),
			).Exec(l.ctx); err != nil {
			return nil, err
		}
	}
	return &types.UpdateWorkspaceWebLinksResponse{
		Message: app.HttpOK,
	}, nil
}
