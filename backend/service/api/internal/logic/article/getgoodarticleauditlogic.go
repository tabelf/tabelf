package article

import (
	"context"
	"tabelf/backend/gen/entschema/goodarticle"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodArticleAuditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodArticleAuditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodArticleAuditLogic {
	return &GetGoodArticleAuditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodArticleAuditLogic) GetGoodArticleAudit(req *types.GetGoodArticleAuditRequest) (resp *types.GetGoodArticleAuditResponse, err error) {
	hasAuthority, err := base.HasAdminAuthority(l.ctx, req.UserUID)
	if err != nil {
		return nil, err
	}
	if !hasAuthority {
		return nil, app.ErrAccountAdminNotExist(l.ctx)
	}
	articles, err := app.EntClient.GoodArticle.Query().Where(
		goodarticle.Status(app.WaitAuditStatus),
		goodarticle.DeactivatedAtIsNil(),
	).All(l.ctx)
	if err != nil {
		return nil, err
	}
	data := make([]*types.GoodArticleRecommend, 0)
	for _, article := range articles {
		data = append(data, &types.GoodArticleRecommend{
			UID:         article.UID,
			Title:       article.Title,
			Description: article.Description,
			Image:       article.Image,
			Source:      article.Source,
			Link:        article.Link,
			Star:        article.Star,
			HasStar:     false,
		})
	}
	return &types.GetGoodArticleAuditResponse{
		Data: data,
	}, nil
}
