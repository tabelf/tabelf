package account

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type GetAuthLoginQrCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAuthLoginQrCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthLoginQrCodeLogic {
	return &GetAuthLoginQrCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuthLoginQrCodeLogic) GetAuthLoginQrCode(req *types.GetAuthLoginQrCodeRequest) (resp *types.GetAuthLoginQrCodeResponse, err error) {
	accessToken, err := base.GetWechatAccessToken(l.ctx)
	if err != nil {
		return nil, err
	}
	ticket, err := base.GetQrExpiredTicket(l.ctx, accessToken, app.Wechat.QrExpired)
	if err != nil {
		return nil, err
	}
	expiredAt := time.Now().Add(time.Second * time.Duration(ticket.ExpireSeconds))
	return &types.GetAuthLoginQrCodeResponse{
		ExpiredAt: app.GetTime(expiredAt),
		URL:       ticket.URL,
	}, nil
}
