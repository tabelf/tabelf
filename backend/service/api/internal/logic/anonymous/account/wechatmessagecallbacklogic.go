package account

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	entaccount "tabelf/backend/gen/entschema/account"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
	"tabelf/backend/spec/schema"
)

type WechatMessageCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWechatMessageCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WechatMessageCallbackLogic {
	return &WechatMessageCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WechatMessageCallbackLogic) WechatMessageCallback(req *types.WechatMessageCallbackRequest) (resp *types.WechatMessageCallbackResponse, err error) {
	switch req.MsgType {
	case app.MsgTextType, app.MsgEventType:
		if req.Content == app.AuthCodeMsg ||
			req.EventKey == app.MenuAuthCodeMsg {
			exist, err := app.EntClient.Account.Query().Where(
				entaccount.WxOpenid(req.FromUserName),
			).Exist(l.ctx)
			if err != nil {
				return nil, err
			}
			var authCode string
			for {
				authCode = fmt.Sprintf("%s%s", time.Now().Format("20060102"), app.RandomString(app.AuthCodeLen))
				hasGen, err := app.EntClient.Account.Query().Where(
					entaccount.AuthCode(authCode),
				).Exist(l.ctx)
				if err != nil {
					return nil, err
				}
				if !hasGen {
					break
				}
			}
			if exist {
				if err := app.EntClient.Account.Update().
					SetAuthCode(authCode).
					SetAuthExpired(time.Now().Add(app.AuthCodeExpiredAt)).
					Where(entaccount.WxOpenid(req.FromUserName)).
					Exec(l.ctx); err != nil {
					return nil, err
				}
			} else {
				if err = app.EntClient.Account.Create().
					SetNickname(app.UnknownUserName + app.RandomString(app.UnknownSuffix)).
					SetSex(app.UnknownSex).
					SetWxOpenid(req.FromUserName).
					SetImage(app.UnknownUserImage).
					SetURLLimit(app.NormalUserURLLimit).
					SetAddress(schema.Address{}).
					SetAuthCode(authCode).
					SetAuthExpired(time.Now().Add(app.AuthCodeExpiredAt)).
					Exec(l.ctx); err != nil {
					return nil, err
				}
			}
			return &types.WechatMessageCallbackResponse{
				Message: fmt.Sprintf(app.WechatReplyTemplate, authCode[8:]),
			}, nil
		}
	}
	// 返回给微信需要要用 "" 空引号
	return &types.WechatMessageCallbackResponse{
		Message: "",
	}, nil
}
