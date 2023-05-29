package article

import (
	"context"
	"tabelf/backend/gen/entschema/goodarticle"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodArticleSubmitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodArticleSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodArticleSubmitLogic {
	return &GetGoodArticleSubmitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodArticleSubmitLogic) GetGoodArticleSubmit(req *types.GetGoodArticleSubmitRequest) (resp *types.GetGoodArticleSubmitResponse, err error) {
	articles, err := app.EntClient.GoodArticle.Query().Where(
		goodarticle.UserUID(req.UserUID),
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
			Status:      article.Status,
		})
	}
	return &types.GetGoodArticleSubmitResponse{
		Data: data,
	}, nil
}
