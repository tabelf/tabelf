package account

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
)

type GetAuthLoginWechatAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAuthLoginWechatAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthLoginWechatAuthorityLogic {
	return &GetAuthLoginWechatAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuthLoginWechatAuthorityLogic) GetAuthLoginWechatAuthority(req *types.GetAuthLoginWechatAuthorityRequest) (resp *types.GetAuthLoginWechatAuthorityResponse, err error) {
	return &types.GetAuthLoginWechatAuthorityResponse{
		Echostr: req.Echostr,
	}, nil
}
