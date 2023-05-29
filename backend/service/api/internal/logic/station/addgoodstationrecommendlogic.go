package station

import (
	"context"
	"net/url"
	"strings"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGoodStationRecommendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddGoodStationRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGoodStationRecommendLogic {
	return &AddGoodStationRecommendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddGoodStationRecommendLogic) AddGoodStationRecommend(req *types.AddGoodStationRecommendRequest) (resp *types.AddGoodStationRecommendResponse, err error) {
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
	if !strings.HasPrefix(req.Link, "http") {
		req.Link = app.HTTPProtoType + req.Link
	}
	info, err := base.GetWebLinkInfo(l.ctx, req.Link)
	if err != nil {
		return nil, err
	}
	stationCreate := app.EntClient.Station.Create().
		SetTitle(info.Title).
		SetDescription(info.Description).
		SetLink(req.Link).
		SetIcon(info.Image).
		SetImage(info.Image).
		SetSource(req.SiteName).
		SetUserUID(req.UserUID).
		SetStatus(app.Show).
		SetCategoryUID(req.CategoryUID).
		SetTags(make([]string, 0))

	if app.IsNotBlank(req.Title) {
		stationCreate.SetTitle(req.Title)
	}
	if app.IsNotBlank(req.Description) {
		stationCreate.SetDescription(req.Description)
	}
	if app.IsNotBlank(req.Image) {
		stationCreate.SetImage(req.Image)
	}
	if len(req.Tags) > 0 {
		stationCreate.SetTags(req.Tags)
	}
	if err = stationCreate.Exec(l.ctx); err != nil {
		return nil, err
	}
	return &types.AddGoodStationRecommendResponse{
		Message: app.HttpOK,
	}, nil
}

func (l *AddGoodStationRecommendLogic) ValidateGoodStationRecommend(req *types.AddGoodStationRecommendRequest) error {
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
