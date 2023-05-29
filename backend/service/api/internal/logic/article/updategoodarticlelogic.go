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

type UpdateGoodArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGoodArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGoodArticleLogic {
	return &UpdateGoodArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGoodArticleLogic) UpdateGoodArticle(req *types.UpdateGoodArticleRequest) (resp *types.UpdateGoodArticleResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	hasAuthority, err := base.HasAdminAuthority(l.ctx, req.UserUID)
	if err != nil {
		return nil, err
	}
	if !hasAuthority {
		return nil, app.ErrAccountAdminNotExist(l.ctx)
	}
	update := app.EntClient.GoodArticle.Update().
		SetStatus(app.WaitAuditStatus)
	if app.IsNotBlank(req.CategoryUID) {
		update.SetCategoryUID(req.CategoryUID)
	}
	if app.IsNotBlank(req.Title) {
		update.SetTitle(req.Title)
	}
	if app.IsNotBlank(req.Source) {
		update.SetSource(req.Source)
	}
	if app.IsNotBlank(req.Image) {
		update.SetImage(req.Image)
	}
	if err = update.Where(
		goodarticle.UID(req.ArticleUID),
		goodarticle.DeactivatedAtIsNil(),
	).Exec(l.ctx); err != nil {
		return nil, err
	}
	article, err := app.EntClient.GoodArticle.Query().Where(
		goodarticle.UID(req.ArticleUID),
		goodarticle.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		return nil, err
	}
	return &types.UpdateGoodArticleResponse{
		UID:         article.UID,
		CategoryUID: article.CategoryUID,
		Title:       article.Title,
		Link:        article.Link,
		Source:      article.Source,
		Image:       article.Image,
		Status:      article.Status,
	}, nil
}
