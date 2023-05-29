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

type GetRecyclingPersonalWebLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRecyclingPersonalWebLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecyclingPersonalWebLinkLogic {
	return &GetRecyclingPersonalWebLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRecyclingPersonalWebLinkLogic) GetRecyclingPersonalWebLink(req *types.GetRecyclingPersonalWebLinkRequest) (resp *types.GetRecyclingPersonalWebLinkResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	links, err := app.EntClient.WebLink.Query().Where(
		entweblink.UserUID(req.UserUID),
		entweblink.ForeverDelete(false),
		entweblink.DeactivatedAtNotNil(),
	).All(l.ctx)
	if err != nil {
		return nil, err
	}
	webLinks := make([]*types.PersonalWebLink, 0)
	for _, link := range links {
		webLinks = append(webLinks, &types.PersonalWebLink{
			LinkUID:     link.UID,
			Title:       link.Title,
			Image:       link.Image,
			Link:        link.Link,
			FileType:    link.FileType,
			Description: link.Description,
			Sequence:    link.Sequence,
		})
	}
	return &types.GetRecyclingPersonalWebLinkResponse{
		WebLinks: webLinks,
	}, nil
}
