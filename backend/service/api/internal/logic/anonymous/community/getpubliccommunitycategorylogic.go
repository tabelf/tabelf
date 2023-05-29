package community

import (
	"context"
	"tabelf/backend/gen/entschema"
	entcommunity "tabelf/backend/gen/entschema/community"
	entcommunitycategory "tabelf/backend/gen/entschema/communitycategory"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublicCommunityCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPublicCommunityCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublicCommunityCategoryLogic {
	return &GetPublicCommunityCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPublicCommunityCategoryLogic) GetPublicCommunityCategory(req *types.GetPublicCommunityCategoryRequest) (resp *types.GetPublicCommunityCategoryResponse, err error) {
	stationCategories, err := app.EntClient.CommunityCategory.Query().Where(
		entcommunitycategory.DeactivatedAtIsNil(),
		entcommunitycategory.Status(app.Show),
	).Order(entschema.Asc(entcommunitycategory.FieldSequence)).
		All(l.ctx)
	if err != nil {
		return nil, err
	}
	categories := make([]*types.PublicCommunityCategory, 0)
	for _, c := range stationCategories {
		categories = append(categories, &types.PublicCommunityCategory{
			UID:  c.UID,
			Name: c.Name,
		})
	}

	hasAdmin, wait := false, 0
	if app.IsNotBlank(req.UserUID) {
		hasAdmin, err = base.HasAdminAuthority(l.ctx, req.UserUID)
		if err != nil {
			return nil, err
		}
		if hasAdmin {
			wait, err = app.EntClient.Community.Query().Where(
				entcommunity.Status(app.WaitAuditStatus),
				entcommunity.DeactivatedAtIsNil(),
			).Count(l.ctx)
			if err != nil {
				return nil, err
			}
		}
	}
	return &types.GetPublicCommunityCategoryResponse{
		Wait:                    wait,
		HasAdmin:                hasAdmin,
		PublicCommunityCategory: categories,
	}, nil
}
