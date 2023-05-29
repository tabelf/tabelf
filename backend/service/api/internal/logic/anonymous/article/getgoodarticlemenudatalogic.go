package article

import (
	"context"
	"tabelf/backend/gen/entschema/goodarticle"
	"tabelf/backend/gen/entschema/goodarticlemeta"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodArticleMenuDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodArticleMenuDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodArticleMenuDataLogic {
	return &GetGoodArticleMenuDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodArticleMenuDataLogic) GetGoodArticleMenuData(req *types.GetGoodArticleMenuDataRequest) (resp *types.GetGoodArticleMenuDataResponse, err error) {
	if app.IsBlank(req.UserUID) {
		return &types.GetGoodArticleMenuDataResponse{
			HasAnonymous: true,
		}, nil
	}
	collection, err := app.EntClient.GoodArticleMeta.Query().Where(
		goodarticlemeta.HasStar(true),
		goodarticlemeta.UserUID(req.UserUID),
		goodarticlemeta.DeactivatedAtIsNil(),
	).Count(l.ctx)
	if err != nil {
		return nil, err
	}
	publish, err := app.EntClient.GoodArticle.Query().Where(
		goodarticle.UserUID(req.UserUID),
		goodarticle.DeactivatedAtIsNil(),
	).Count(l.ctx)
	if err != nil {
		return nil, err
	}
	hasAuthority, err := base.HasAdminAuthority(l.ctx, req.UserUID)
	if err != nil {
		return nil, err
	}
	if !hasAuthority {
		return &types.GetGoodArticleMenuDataResponse{
			HasAnonymous: false,
			HasAuthority: hasAuthority,
			Collection:   collection,
			Audit:        0,
			Publish:      publish,
		}, nil
	}
	audit, err := app.EntClient.GoodArticle.Query().Where(
		goodarticle.Status(app.WaitAuditStatus),
		goodarticle.DeactivatedAtIsNil(),
	).Count(l.ctx)
	if err != nil {
		return nil, err
	}
	return &types.GetGoodArticleMenuDataResponse{
		HasAnonymous: false,
		HasAuthority: hasAuthority,
		Collection:   collection,
		Audit:        audit,
		Publish:      publish,
	}, nil
}
