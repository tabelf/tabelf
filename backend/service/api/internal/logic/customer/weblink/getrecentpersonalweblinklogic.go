package weblink

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"tabelf/backend/gen/entschema"
	entweblink "tabelf/backend/gen/entschema/weblink"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type GetRecentPersonalWebLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRecentPersonalWebLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecentPersonalWebLinkLogic {
	return &GetRecentPersonalWebLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRecentPersonalWebLinkLogic) GetRecentPersonalWebLink(req *types.GetRecentPersonalWebLinkRequest) (resp *types.GetRecentPersonalWebLinkResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	links, err := app.EntClient.WebLink.Query().Where(
		entweblink.UserUID(req.UserUID),
		entweblink.DeactivatedAtIsNil(),
	).Order(entschema.Desc(entweblink.FieldUpdatedAt)).
		Offset(0).Limit(app.RecentWebLinkMaxLimitNum).
		All(l.ctx)
	if err != nil {
		return nil, err
	}
	webLinks := make([]*types.RecentWebLink, 0)
	if len(links) == 0 {
		return &types.GetRecentPersonalWebLinkResponse{
			WebLinks: webLinks,
		}, nil
	}
	for _, link := range links {
		webLinks = append(webLinks, &types.RecentWebLink{
			LinkUID:     link.UID,
			Title:       link.Title,
			Image:       link.Image,
			Link:        link.Link,
			FileType:    link.FileType,
			Description: link.Description,
			UpdatedAt:   app.GetTime(link.UpdatedAt),
		})
	}
	return &types.GetRecentPersonalWebLinkResponse{
		WebLinks: webLinks,
	}, nil
}
