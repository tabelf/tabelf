package link

import (
	"context"
	"fmt"
	"tabelf/backend/common"
	"tabelf/backend/gen/entschema"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"
	"time"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareFolderLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareFolderLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareFolderLinkLogic {
	return &ShareFolderLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareFolderLinkLogic) ShareFolderLink(req *types.ShareFolderLinkRequest) (resp *types.ShareFolderLinkResponse, err error) {
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
	folder, err := app.EntClient.PersonalFolder.Query().Where(
		entpersonalfolder.UID(req.FolderUID),
		entpersonalfolder.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		if entschema.IsNotFound(err) {
			return nil, app.ErrPersonalFolderNotExist(l.ctx)
		}
		return nil, err
	}
	expiredAt := time.Now().AddDate(0, 0, req.ExpiredDay)
	if req.ExpiredDay == -1 {
		expiredAt = app.ParseTime("2099-12-31 23:59:59")
	}
	shareUID := common.NewUUID()
	if err = app.EntClient.ShareLink.Create().
		SetUID(shareUID).
		SetFolderUID(folder.UID).
		SetUserUID(req.UserUID).
		SetAuthority(req.Authority).
		SetExpiredAt(expiredAt).
		Exec(l.ctx); err != nil {
		return nil, err
	}
	return &types.ShareFolderLinkResponse{
		ShareUID:   shareUID,
		ShareLink:  fmt.Sprintf("%s/v/%s", app.Basic.Domain, shareUID),
		Authority:  req.Authority,
		ExpiredDay: req.ExpiredDay,
	}, nil
}
