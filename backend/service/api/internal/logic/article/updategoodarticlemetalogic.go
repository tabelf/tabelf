package article

import (
	"context"
	"tabelf/backend/service/api/internal/logic/base"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGoodArticleMetaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGoodArticleMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGoodArticleMetaLogic {
	return &UpdateGoodArticleMetaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGoodArticleMetaLogic) UpdateGoodArticleMeta(req *types.UpdateGoodArticleMetaRequest) (resp *types.UpdateGoodArticleMetaResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	return base.UpdateGoodArticleMeta(l.ctx, req.ArticleUID, req.UserUID, req.MetaType)
}
