package global

import (
	"context"
	"fmt"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SayHelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSayHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHelloLogic {
	return &SayHelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SayHelloLogic) SayHello(req *types.HelloRequest) (resp *types.HelloResponse, err error) {
	return &types.HelloResponse{
		Message: fmt.Sprintf("Hello %s", req.Name),
	}, nil
}
