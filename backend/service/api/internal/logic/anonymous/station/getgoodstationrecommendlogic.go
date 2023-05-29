package station

import (
	"context"
	"tabelf/backend/gen/entschema"
	entstation "tabelf/backend/gen/entschema/station"
	entstationcategory "tabelf/backend/gen/entschema/stationcategory"
	entstationmeta "tabelf/backend/gen/entschema/stationmeta"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodStationRecommendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodStationRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodStationRecommendLogic {
	return &GetGoodStationRecommendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodStationRecommendLogic) GetGoodStationRecommend(req *types.GetGoodStationRecommendRequest) (resp *types.GetGoodStationRecommendResponse, err error) {
	stationQuery := app.EntClient.Station.Query().Where(
		entstation.Status(app.Show),
		entstation.DeactivatedAtIsNil(),
	)
	if app.IsNotBlank(req.CategoryUID) {
		stationQuery.Where(entstation.CategoryUID(req.CategoryUID))
	}
	switch req.Sorted {
	case app.StationView: // 全部: 查看
		stationQuery.Order(entschema.Desc(entstation.FieldView))
	case app.StationPraise: // 点赞最多
		stationQuery.Order(entschema.Desc(entstation.FieldPraise))
	case app.StationStar: // 收藏最多
		stationQuery.Order(entschema.Desc(entstation.FieldStar))
	case app.StationNew: // 最新
		stationQuery.Order(entschema.Desc(entstation.FieldCreatedAt))
	default:
		return nil, app.ErrStationSortedInvalid(l.ctx)
	}
	count, err := stationQuery.Count(l.ctx)
	if err != nil {
		return nil, err
	}
	data := make([]*types.GoodStationRecommend, 0)
	if count != 0 {
		stations, err := stationQuery.Offset(req.Offset).Limit(req.Limit).All(l.ctx)
		if err != nil {
			return nil, err
		}
		categoryUIDs := make([]string, 0)
		metaUIDs := make([]string, 0)
		for _, station := range stations {
			categoryUIDs = append(categoryUIDs, station.CategoryUID)
			metaUIDs = append(metaUIDs, station.UID)
		}
		categories, err := app.EntClient.StationCategory.Query().Where(
			entstationcategory.UIDIn(categoryUIDs...),
			entstationcategory.DeactivatedAtIsNil(),
		).All(l.ctx)
		categoryMap := make(map[string]*entschema.StationCategory)
		for _, category := range categories {
			categoryMap[category.UID] = category
		}
		var metas []*entschema.StationMeta
		if app.IsNotBlank(req.UserUID) {
			metas, err = app.EntClient.StationMeta.Query().Where(
				entstationmeta.StationUIDIn(metaUIDs...),
				entstationmeta.UserUID(req.UserUID),
				entstationmeta.DeactivatedAtIsNil(),
			).All(l.ctx)
			if err != nil {
				return nil, err
			}
		}
		metaMap := make(map[string]*entschema.StationMeta)
		for _, meta := range metas {
			metaMap[meta.StationUID] = meta
		}
		for _, station := range stations {
			name := ""
			if c, ok := categoryMap[station.CategoryUID]; ok {
				name = c.Name
			}
			hasPraise, hasStar := false, false
			if m, ok := metaMap[station.UID]; ok {
				hasPraise = m.HasPraise
				hasStar = m.HasStar
			}
			data = append(data, &types.GoodStationRecommend{
				UID:          station.UID,
				Title:        station.Title,
				Description:  station.Description,
				Image:        station.Image,
				Tags:         station.Tags,
				Icon:         station.Icon,
				Source:       station.Source,
				Link:         station.Link,
				Praise:       station.Praise,
				HasPraise:    hasPraise,
				Star:         station.Star,
				HasStar:      hasStar,
				View:         station.View,
				CategoryName: name,
			})
		}
	}
	return &types.GetGoodStationRecommendResponse{
		Total: count,
		Data:  data,
	}, nil
}
