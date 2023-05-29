package message

import (
	"context"
	entmessage "tabelf/backend/gen/entschema/message"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"
	"time"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelMessageLogic {
	return &DelMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelMessageLogic) DelMessage(req *types.DelMessageRequest) (resp *types.DelMessageResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	if err = app.EntClient.Message.Update().
		SetDeactivatedAt(time.Now()).
		Where(
			entmessage.UID(req.MessageUID),
			entmessage.UserUID(req.UserUID),
			entmessage.DeactivatedAtIsNil(),
		).
		Exec(l.ctx); err != nil {
		return nil, err
	}
	return &types.DelMessageResponse{
		Message: app.HttpOK,
	}, nil
}
