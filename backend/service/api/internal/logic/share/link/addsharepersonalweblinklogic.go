package link

import (
	"context"
	"tabelf/backend/gen/entschema"
	entsharelink "tabelf/backend/gen/entschema/sharelink"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSharePersonalWebLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSharePersonalWebLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSharePersonalWebLinkLogic {
	return &AddSharePersonalWebLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSharePersonalWebLinkLogic) AddSharePersonalWebLink(req *types.AddSharePersonalWebLinkRequest) (resp *types.AddSharePersonalWebLinkResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	if app.IsBlank(req.URL) {
		return nil, app.ErrCustomerWebLinkEmpty(l.ctx)
	}
	// 查询该文件的协作链接
	shareLink, err := app.EntClient.ShareLink.Query().Where(
		entsharelink.UID(req.ShareUID),
		entsharelink.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		if entschema.IsNotFound(err) {
			return nil, app.ErrPersonalFolderShareLinkNotExist(l.ctx)
		}
		return nil, err
	}
	if err = base.AddWebLink(l.ctx, req.URL, req.Title, shareLink.UserUID, req.WorkspaceUID); err != nil {
		return nil, err
	}
	return &types.AddSharePersonalWebLinkResponse{
		Message: app.HttpOK,
	}, nil
}
