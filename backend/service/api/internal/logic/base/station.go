package base

import (
	"context"
	"tabelf/backend/gen/entschema"
	entstation "tabelf/backend/gen/entschema/station"
	entstationmeta "tabelf/backend/gen/entschema/stationmeta"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

func UpdateGoodStationMeta(ctx context.Context, stationUID, userUID string, metaType int) (resp *types.UpdateGoodStationMetaResponse, err error) {
	station, err := app.EntClient.Station.Query().Where(
		entstation.UID(stationUID),
		entstation.DeactivatedAtIsNil(),
	).First(ctx)
	if err != nil {
		return nil, err
	}
	meta, err := app.EntClient.StationMeta.Query().Where(
		entstationmeta.UserUID(userUID),
		entstationmeta.StationUID(stationUID),
		entstationmeta.DeactivatedAtIsNil(),
	).First(ctx)
	if err != nil && !entschema.IsNotFound(err) {
		return nil, err
	}
	if err = app.WithTx(ctx, app.EntClient, func(tx *entschema.Tx) (err error) {
		stationUpdate := tx.Station.Update()
		if meta == nil {
			metaCreate := tx.StationMeta.Create().
				SetUserUID(userUID).
				SetStationUID(stationUID)
			switch metaType {
			case app.StationView: // 全部: 查看
				metaCreate.SetHasView(true)
				stationUpdate.AddView(1)
			case app.StationPraise: // 点赞最多
				metaCreate.SetHasPraise(true)
				stationUpdate.AddPraise(1)
			case app.StationStar: // 收藏最多
				metaCreate.SetHasStar(true)
				stationUpdate.AddStar(1)
			default:
				return app.ErrStationSortedInvalid(ctx)
			}
			if err = metaCreate.Exec(ctx); err != nil {
				return err
			}
		} else {
			metaUpdate := tx.StationMeta.Update()
			switch metaType {
			case app.StationView: // 查看
				if !meta.HasView {
					metaUpdate.SetHasView(true)
					stationUpdate.AddView(1)
				}
			case app.StationPraise: // 点赞最多
				if meta.HasPraise {
					metaUpdate.SetHasPraise(false)
					stationUpdate.AddPraise(-1)
				} else {
					metaUpdate.SetHasPraise(true)
					stationUpdate.AddPraise(1)
				}
			case app.StationStar: // 收藏最多
				if !meta.HasStar {
					metaUpdate.SetHasStar(true)
					stationUpdate.AddStar(1)
				}
			default:
				return app.ErrStationSortedInvalid(ctx)
			}
			if err = metaUpdate.
				Where(entstationmeta.ID(meta.ID)).
				Exec(ctx); err != nil {
				return err
			}
		}
		if err = stationUpdate.
			Where(entstation.ID(station.ID)).
			Exec(ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	station, err = app.EntClient.Station.Query().Where(
		entstation.ID(station.ID),
	).First(ctx)
	if err != nil {
		return nil, err
	}
	meta, err = app.EntClient.StationMeta.Query().Where(
		entstationmeta.UserUID(userUID),
		entstationmeta.StationUID(stationUID),
		entstationmeta.DeactivatedAtIsNil(),
	).First(ctx)
	return &types.UpdateGoodStationMetaResponse{
		UID:       station.UID,
		Praise:    station.Praise,
		HasPraise: meta.HasPraise,
		Star:      station.Star,
		HasStar:   meta.HasStar,
		View:      station.View,
	}, nil
}
