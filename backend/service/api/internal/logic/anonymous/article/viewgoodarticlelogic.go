package article

import (
	"context"
	"tabelf/backend/gen/entschema/goodarticle"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewGoodArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewViewGoodArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewGoodArticleLogic {
	return &ViewGoodArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewGoodArticleLogic) ViewGoodArticle(req *types.ViewGoodArticleRequest) (resp *types.ViewGoodArticleResponse, err error) {
	if err = app.EntClient.GoodArticle.Update().AddView(1).Where(
		goodarticle.UID(req.ArticleUID),
		goodarticle.DeactivatedAtIsNil(),
	).Exec(l.ctx); err != nil {
		return nil, err
	}
	return &types.ViewGoodArticleResponse{
		Message: app.HttpOK,
	}, nil
}
