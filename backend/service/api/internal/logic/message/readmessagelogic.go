package message

import (
	"context"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/api/models"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReadMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReadMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadMessageLogic {
	return &ReadMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReadMessageLogic) ReadMessage(req *types.ReadMessageRequest) (resp *types.ReadMessageResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	if err = models.ReadMessage(l.ctx, req.MessageUID, req.UserUID); err != nil {
		return nil, err
	}
	return &types.ReadMessageResponse{
		Message: app.HttpOK,
	}, nil
}
