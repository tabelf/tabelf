package station

import (
	"context"
	"net/url"
	entstation "tabelf/backend/gen/entschema/station"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGoodStationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGoodStationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGoodStationLogic {
	return &UpdateGoodStationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGoodStationLogic) UpdateGoodStation(req *types.UpdateGoodStationRequest) (resp *types.UpdateGoodStationResponse, err error) {
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
	if err = l.ValidateGoodStationRecommend(req); err != nil {
		return nil, err
	}
	station, err := app.EntClient.Station.Query().Where(
		entstation.UID(req.StationUID),
		entstation.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		return nil, err
	}
	update := app.EntClient.Station.Update().
		SetTitle(req.Title).
		SetDescription(req.Description).
		SetLink(req.Link).
		SetIcon(req.Image).
		SetImage(req.Image).
		SetSource(req.SiteName).
		SetUserUID(req.UserUID).
		SetCategoryUID(req.CategoryUID).
		SetTags(req.Tags)
	if station.Link != req.Link {
		info, err := base.GetWebLinkInfo(l.ctx, req.Link)
		if err != nil {
			return nil, err
		}
		update = update.SetTitle(info.Title).
			SetDescription(info.Description).
			SetIcon(info.Image).
			SetImage(info.Image)
	}
	if err = update.Where(entstation.ID(station.ID)).
		Exec(l.ctx); err != nil {
		return nil, err
	}
	return &types.UpdateGoodStationResponse{
		Message: app.HttpOK,
	}, nil
}

func (l *UpdateGoodStationLogic) ValidateGoodStationRecommend(req *types.UpdateGoodStationRequest) error {
	if app.IsBlank(req.CategoryUID) {
		return app.ErrStationCategoryEmpty(l.ctx)
	}
	if app.IsBlank(req.Link) {
		return app.ErrStationLinkEmpty(l.ctx)
	}
	_, err := url.Parse(req.Link)
	if err != nil {
		return app.ErrStationLinkInvalid(l.ctx)
	}
	return nil
}
