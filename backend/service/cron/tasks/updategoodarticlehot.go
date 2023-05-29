package tasks

import (
	"sort"
	"time"

	"tabelf/backend/gen/entschema"
	"tabelf/backend/gen/entschema/goodarticle"
	"tabelf/backend/gen/entschema/goodarticlehot"
	"tabelf/backend/service/app"
)

// UpdateGoodArticleHot 每天查询最新热度文章.
func UpdateGoodArticleHot(jobCtx JobContext, config app.Config) {
	ctx := jobCtx.Context
	current := time.Now()

	hots, err := app.EntClient.GoodArticleHot.Query().Where(
		goodarticlehot.HasExpired(false),
		goodarticlehot.DeactivatedAtIsNil(),
	).All(ctx)
	if err != nil {
		return
	}
	// 查询8天以内的好文推荐列表
	articles, err := app.EntClient.GoodArticle.Query().Where(
		goodarticle.CreatedAtGTE(current.Add(-8*24*time.Hour)),
		goodarticle.Status(app.PassStatus),
		goodarticle.DeactivatedAtIsNil(),
	).All(ctx)
	if err != nil {
		return
	}
	if len(articles) < app.HotArticleNum {
		articleUID := make([]string, 0)
		for _, article := range articles {
			articleUID = append(articleUID, article.UID)
		}
		goodArticles, err := app.EntClient.GoodArticle.Query().Where(
			goodarticle.UIDNotIn(articleUID...),
			goodarticle.Status(app.PassStatus),
			goodarticle.DeactivatedAtIsNil(),
		).Order(entschema.Desc(goodarticle.FieldCreatedAt)).
			All(ctx)
		if err != nil {
			return
		}
		articles = append(articles, goodArticles...)
	}
	if len(articles) == 0 {
		return
	}
	type WeightArticle struct {
		weight  float64
		article *entschema.GoodArticle
	}
	weights := make([]*WeightArticle, 0)
	for _, article := range articles {
		weight := app.HotArticleWeight(article)
		weights = append(weights, &WeightArticle{
			weight:  weight,
			article: article,
		})
	}
	sort.Slice(weights, func(i, j int) bool {
		return weights[i].weight > weights[j].weight
	})
	expiredUIDs := make([]string, 0)
	for _, hot := range hots {
		expiredUIDs = append(expiredUIDs, hot.UID)
	}
	if err = app.WithTx(ctx, app.EntClient, func(tx *entschema.Tx) (err error) {
		if len(expiredUIDs) != 0 {
			if err = tx.GoodArticleHot.Update().SetHasExpired(true).Where(
				goodarticlehot.UIDIn(expiredUIDs...),
			).Exec(ctx); err != nil {
				return err
			}
		}
		for _, weight := range weights {
			if err = tx.GoodArticleHot.Create().
				SetArticleUID(weight.article.UID).
				Exec(ctx); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return
	}
}
