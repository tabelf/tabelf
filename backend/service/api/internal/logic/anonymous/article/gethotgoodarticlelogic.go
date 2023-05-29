package article

import (
	"context"
	"tabelf/backend/gen/entschema"
	"tabelf/backend/gen/entschema/goodarticle"
	"tabelf/backend/gen/entschema/goodarticlehot"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHotGoodArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHotGoodArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHotGoodArticleLogic {
	return &GetHotGoodArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHotGoodArticleLogic) GetHotGoodArticle(req *types.GetHotGoodArticleRequest) (resp *types.GetHotGoodArticleResponse, err error) {
	hots, err := app.EntClient.GoodArticleHot.Query().Where(
		goodarticlehot.HasExpired(false),
		goodarticlehot.DeactivatedAtIsNil(),
	).All(l.ctx)
	if err != nil {
		return
	}
	articles := make([]*types.HotGoodArticle, 0)
	for _, hot := range hots {
		article, err := app.EntClient.GoodArticle.Query().Where(
			goodarticle.UID(hot.ArticleUID),
			goodarticle.DeactivatedAtIsNil(),
		).First(l.ctx)
		if err != nil && !entschema.IsNotFound(err) {
			return nil, err
		}
		if articles == nil {
			continue
		}
		articles = append(articles, &types.HotGoodArticle{
			UID:   article.UID,
			Title: article.Title,
			Link:  article.Link,
		})
	}
	return &types.GetHotGoodArticleResponse{
		Data: articles,
	}, nil
}
