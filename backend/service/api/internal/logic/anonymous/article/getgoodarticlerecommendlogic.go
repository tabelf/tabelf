package article

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"tabelf/backend/gen/entschema"
	entgoodarticle "tabelf/backend/gen/entschema/goodarticle"
	entgoodarticlemeta "tabelf/backend/gen/entschema/goodarticlemeta"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type GetGoodArticleRecommendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodArticleRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodArticleRecommendLogic {
	return &GetGoodArticleRecommendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodArticleRecommendLogic) GetGoodArticleRecommend(req *types.GetGoodArticleRecommendRequest) (resp *types.GetGoodArticleRecommendResponse, err error) {
	articleQuery := app.EntClient.GoodArticle.Query().Where(
		entgoodarticle.Status(app.PassStatus),
		entgoodarticle.DeactivatedAtIsNil(),
	)
	if app.IsNotBlank(req.CategoryUID) {
		articleQuery.Where(entgoodarticle.CategoryUID(req.CategoryUID))
	}
	switch req.Sorted {
	case app.ArticleView: // 全部: 查看
		articleQuery.Order(entschema.Desc(entgoodarticle.FieldView))
	case app.ArticleNew: // 最新
		articleQuery.Order(entschema.Desc(entgoodarticle.FieldCreatedAt))
	default:
		return nil, app.ErrArticleSortedInvalid(l.ctx)
	}
	count, err := articleQuery.Count(l.ctx)
	if err != nil {
		return nil, err
	}
	data := make([]*types.GoodArticleRecommend, 0)
	if count != 0 {
		articles, err := articleQuery.Offset(req.Offset).Limit(req.Limit).All(l.ctx)
		if err != nil {
			return nil, err
		}
		metaUIDs := make([]string, 0)
		for _, article := range articles {
			metaUIDs = append(metaUIDs, article.UID)
		}
		var metas []*entschema.GoodArticleMeta
		if app.IsNotBlank(req.UserUID) {
			metas, err = app.EntClient.GoodArticleMeta.Query().Where(
				entgoodarticlemeta.ArticleUIDIn(metaUIDs...),
				entgoodarticlemeta.UserUID(req.UserUID),
				entgoodarticlemeta.DeactivatedAtIsNil(),
			).All(l.ctx)
			if err != nil {
				return nil, err
			}
		}
		metaMap := make(map[string]*entschema.GoodArticleMeta)
		for _, meta := range metas {
			metaMap[meta.ArticleUID] = meta
		}
		for _, article := range articles {
			hasStar := false
			if m, ok := metaMap[article.UID]; ok {
				hasStar = m.HasStar
			}
			data = append(data, &types.GoodArticleRecommend{
				UID:         article.UID,
				Title:       article.Title,
				Description: article.Description,
				Image:       article.Image,
				Source:      article.Source,
				Link:        article.Link,
				Star:        article.Star,
				HasStar:     hasStar,
			})
		}
	}
	return &types.GetGoodArticleRecommendResponse{
		Total: count,
		Data:  data,
	}, nil
}
