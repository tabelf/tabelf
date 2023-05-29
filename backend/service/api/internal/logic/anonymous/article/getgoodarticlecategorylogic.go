package article

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"tabelf/backend/gen/entschema"
	entgoodarticlecategory "tabelf/backend/gen/entschema/goodarticlecategory"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type GetGoodArticleCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodArticleCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodArticleCategoryLogic {
	return &GetGoodArticleCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodArticleCategoryLogic) GetGoodArticleCategory(req *types.GetGoodArticleCategoryRequest) (resp *types.GetGoodArticleCategoryResponse, err error) {
	articleCategories, err := app.EntClient.GoodArticleCategory.Query().Where(
		entgoodarticlecategory.DeactivatedAtIsNil(),
		entgoodarticlecategory.Status(app.Show),
	).Order(entschema.Asc(entgoodarticlecategory.FieldSequence)).
		All(l.ctx)
	if err != nil {
		return nil, err
	}
	categories := make([]*types.GoodArticleCategory, 0)
	for _, c := range articleCategories {
		categories = append(categories, &types.GoodArticleCategory{
			UID:  c.UID,
			Name: c.Name,
		})
	}
	return &types.GetGoodArticleCategoryResponse{
		GoodArticleCategories: categories,
	}, nil
}
