package link

import (
	"context"
	entsharelink "tabelf/backend/gen/entschema/sharelink"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"
	"time"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateShareLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateShareLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateShareLinkLogic {
	return &UpdateShareLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateShareLinkLogic) UpdateShareLink(req *types.UpdateShareLinkRequest) (resp *types.UpdateShareLinkResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	if !app.StringSliceContains([]string{
		app.ShareReadAuthority,
		app.ShareEditAuthority,
	}, req.Authority) {
		return nil, app.ErrCustomerFolderShareAuthorityInvalid(l.ctx)
	}
	if !app.IntSliceContains([]int{-1, 7, 30}, req.ExpiredDay) {
		return nil, app.ErrCustomerFolderShareExpiredDayInvalid(l.ctx)
	}
	expiredAt := time.Now().AddDate(0, 0, req.ExpiredDay)
	if req.ExpiredDay == -1 {
		expiredAt = app.ParseTime("2099-12-31 23:59:59")
	}
	if err = app.EntClient.ShareLink.Update().
		SetAuthority(req.Authority).
		SetValidDay(req.ExpiredDay).
		SetExpiredAt(expiredAt).
		Where(
			entsharelink.UID(req.ShareUID),
			entsharelink.UserUID(req.UserUID),
			entsharelink.FolderUID(req.FolderUID),
			entsharelink.DeactivatedAtIsNil(),
		).Exec(l.ctx); err != nil {
		return nil, err
	}
	return &types.UpdateShareLinkResponse{
		Message: app.HttpOK,
	}, nil
}
