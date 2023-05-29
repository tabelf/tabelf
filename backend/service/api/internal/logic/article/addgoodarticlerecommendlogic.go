package article

import (
	"context"
	"strings"
	"tabelf/backend/gen/entschema/goodarticle"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGoodArticleRecommendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddGoodArticleRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGoodArticleRecommendLogic {
	return &AddGoodArticleRecommendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddGoodArticleRecommendLogic) AddGoodArticleRecommend(req *types.AddGoodArticleRecommendRequest) (resp *types.AddGoodArticleRecommendResponse, err error) {
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
	exist, err := app.EntClient.GoodArticle.Query().Where(
		goodarticle.Link(req.Link),
		goodarticle.DeactivatedAtIsNil(),
	).Exist(l.ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, app.ErrArticleExist(l.ctx)
	}
	if !strings.HasPrefix(req.Link, "http") {
		req.Link = app.HTTPProtoType + req.Link
	}
	info, err := base.GetWebLinkInfo(l.ctx, req.Link)
	if err != nil {
		return nil, err
	}
	articleCreate := app.EntClient.GoodArticle.Create().
		SetTitle(info.Title).
		SetDescription(info.Title).
		SetImage(req.Image).
		SetIcon(info.Image).
		SetSource(req.Source).
		SetLink(req.Link).
		SetUserUID(req.UserUID).
		SetStatus(app.WaitAuditStatus).
		SetCategoryUID(req.CategoryUID)

	if app.IsNotBlank(req.Title) {
		articleCreate.SetTitle(req.Title)
	}
	if err = articleCreate.Exec(l.ctx); err != nil {
		return nil, err
	}
	return &types.AddGoodArticleRecommendResponse{
		Message: app.HttpOK,
	}, nil
}
