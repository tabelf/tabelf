package message

import (
	"context"
	entmessage "tabelf/backend/gen/entschema/message"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReadAllMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReadAllMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadAllMessageLogic {
	return &ReadAllMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReadAllMessageLogic) ReadAllMessage(req *types.ReadAllMessageRequest) (resp *types.ReadAllMessageResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	if err = app.EntClient.Message.Update().
		SetHasRead(true).
		Where(entmessage.UserUID(req.UserUID), entmessage.DeactivatedAtIsNil()).
		Exec(l.ctx); err != nil {
		return nil, err
	}
	return &types.ReadAllMessageResponse{
		Message: app.HttpOK,
	}, nil
}
