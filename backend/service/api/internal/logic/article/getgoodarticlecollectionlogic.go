package article

import (
	"context"
	"tabelf/backend/gen/entschema/goodarticle"
	"tabelf/backend/gen/entschema/goodarticlemeta"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodArticleCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodArticleCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodArticleCollectionLogic {
	return &GetGoodArticleCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodArticleCollectionLogic) GetGoodArticleCollection(req *types.GetGoodArticleCollectionRequest) (resp *types.GetGoodArticleCollectionResponse, err error) {
	metas, err := app.EntClient.GoodArticleMeta.Query().Where(
		goodarticlemeta.HasStar(true),
		goodarticlemeta.UserUID(req.UserUID),
		goodarticlemeta.DeactivatedAtIsNil(),
	).All(l.ctx)
	if err != nil {
		return nil, err
	}
	data := make([]*types.GoodArticleRecommend, 0)
	if len(metas) == 0 {
		return &types.GetGoodArticleCollectionResponse{
			Data: data,
		}, nil
	}
	articleUIDs := make([]string, 0)
	for _, meta := range metas {
		articleUIDs = append(articleUIDs, meta.ArticleUID)
	}
	articles, err := app.EntClient.GoodArticle.Query().Where(
		goodarticle.UIDIn(articleUIDs...),
		goodarticle.DeactivatedAtIsNil(),
	).All(l.ctx)
	if err != nil {
		return nil, err
	}
	for _, article := range articles {
		data = append(data, &types.GoodArticleRecommend{
			UID:         article.UID,
			Title:       article.Title,
			Description: article.Description,
			Image:       article.Image,
			Source:      article.Source,
			Link:        article.Link,
			Star:        article.Star,
			HasStar:     true,
		})
	}
	return &types.GetGoodArticleCollectionResponse{
		Data: data,
	}, nil
}
