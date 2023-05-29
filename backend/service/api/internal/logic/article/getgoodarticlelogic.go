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

type GetGoodArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodArticleLogic {
	return &GetGoodArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodArticleLogic) GetGoodArticle(req *types.GetGoodArticleRequest) (resp *types.GetGoodArticleResponse, err error) {
	hasAuthority, err := base.HasAdminAuthority(l.ctx, req.UserUID)
	if err != nil {
		return nil, err
	}
	if !hasAuthority {
		return nil, app.ErrAccountAdminNotExist(l.ctx)
	}
	article, err := app.EntClient.GoodArticle.Query().Where(
		goodarticle.UID(req.ArticleUID),
		goodarticle.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		return nil, err
	}
	return &types.GetGoodArticleResponse{
		ArticleUID:  article.UID,
		CategoryUID: article.CategoryUID,
		Title:       article.Title,
		Link:        article.Link,
		Source:      article.Source,
		Image:       article.Image,
	}, nil
}
