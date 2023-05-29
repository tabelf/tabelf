package weblink

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	entweblink "tabelf/backend/gen/entschema/weblink"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type DeleteForeverPersonalWebLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteForeverPersonalWebLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteForeverPersonalWebLinkLogic {
	return &DeleteForeverPersonalWebLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteForeverPersonalWebLinkLogic) DeleteForeverPersonalWebLink(req *types.DeleteForeverPersonalWebLinkRequest) (resp *types.DeleteForeverPersonalWebLinkResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	if err = app.WithTx(l.ctx, app.EntClient, func(tx *entschema.Tx) (err error) {
		if err = tx.WebLink.Update().
			SetForeverDelete(true).
			Where(entweblink.UID(req.LinkUID)).
			Exec(l.ctx); err != nil {
			return err
		}
		if err = tx.Account.Update().
			AddURLCount(-1).
			Where(entaccount.UID(req.UserUID)).
			Exec(l.ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &types.DeleteForeverPersonalWebLinkResponse{
		Message: app.HttpOK,
	}, nil
}
