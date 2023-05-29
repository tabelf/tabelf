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

type DelAllMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelAllMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAllMessageLogic {
	return &DelAllMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelAllMessageLogic) DelAllMessage(req *types.DelAllMessageRequest) (resp *types.DelAllMessageResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	if err = app.EntClient.Message.Update().
		SetDeactivatedAt(time.Now()).
		Where(entmessage.UserUID(req.UserUID), entmessage.DeactivatedAtIsNil()).
		Exec(l.ctx); err != nil {
		return nil, err
	}
	return &types.DelAllMessageResponse{
		Message: app.HttpOK,
	}, nil
}
