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

type UpdateGoodArticleAuditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGoodArticleAuditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGoodArticleAuditLogic {
	return &UpdateGoodArticleAuditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGoodArticleAuditLogic) UpdateGoodArticleAudit(req *types.UpdateGoodArticleAuditRequest) (resp *types.UpdateGoodArticleAuditResponse, err error) {
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
	status := app.Then(req.Status, app.PassStatus, app.FailPassStatus).(string)
	if err = app.EntClient.GoodArticle.Update().SetStatus(status).
		Where(goodarticle.UID(req.ArticleUID)).
		Exec(l.ctx); err != nil {
		return nil, err
	}
	return &types.UpdateGoodArticleAuditResponse{
		Message: app.HttpOK,
	}, nil
}
