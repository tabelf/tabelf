package base

import (
	"context"
	"tabelf/backend/gen/entschema"
	"tabelf/backend/gen/entschema/goodarticle"
	"tabelf/backend/gen/entschema/goodarticlemeta"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

func UpdateGoodArticleMeta(ctx context.Context, articleUID, userUID string, metaType int) (resp *types.UpdateGoodArticleMetaResponse, err error) {
	article, err := app.EntClient.GoodArticle.Query().Where(
		goodarticle.UID(articleUID),
		goodarticle.DeactivatedAtIsNil(),
	).First(ctx)
	if err != nil {
		return nil, err
	}
	meta, err := app.EntClient.GoodArticleMeta.Query().Where(
		goodarticlemeta.UserUID(userUID),
		goodarticlemeta.ArticleUID(articleUID),
		goodarticlemeta.DeactivatedAtIsNil(),
	).First(ctx)
	if err != nil && !entschema.IsNotFound(err) {
		return nil, err
	}
	hasStar := false
	if err = app.WithTx(ctx, app.EntClient, func(tx *entschema.Tx) (err error) {
		articleUpdate := tx.GoodArticle.Update()
		if meta == nil {
			metaCreate := tx.GoodArticleMeta.Create().
				SetUserUID(userUID).
				SetArticleUID(articleUID)
			switch metaType {
			case app.ArticleView: // 全部: 查看
				metaCreate.SetHasView(true)
				articleUpdate.AddView(1)
			case app.ArticleUsed: // 使用
				metaCreate.SetHasUsed(true)
				articleUpdate.AddUsed(1)
			case app.ArticleStar: // 收藏最多
				metaCreate.SetHasStar(true)
				articleUpdate.AddStar(1)
				hasStar = true
			default:
				return app.ErrArticleSortedInvalid(ctx)
			}
			if err = metaCreate.Exec(ctx); err != nil {
				return err
			}
		} else {
			metaUpdate := tx.GoodArticleMeta.Update()
			switch metaType {
			case app.StationView: // 查看
				if !meta.HasView {
					metaUpdate.SetHasView(true)
					articleUpdate.AddView(1)
				}
			case app.ArticleUsed: // 使用最多
				if !meta.HasUsed {
					metaUpdate.SetHasUsed(true)
					articleUpdate.AddUsed(1)
				}
			case app.StationStar: // 收藏最多
				if !meta.HasStar {
					metaUpdate.SetHasStar(true)
					articleUpdate.AddStar(1)
					hasStar = true
				} else {
					metaUpdate.SetHasStar(false)
					articleUpdate.AddStar(-1)
				}
			default:
				return app.ErrArticleSortedInvalid(ctx)
			}
			if err = metaUpdate.
				Where(goodarticlemeta.ID(meta.ID)).
				Exec(ctx); err != nil {
				return err
			}
		}
		if err = articleUpdate.
			Where(goodarticle.ID(article.ID)).
			Exec(ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &types.UpdateGoodArticleMetaResponse{
		UID:     article.UID,
		HasStar: hasStar,
	}, nil
}
