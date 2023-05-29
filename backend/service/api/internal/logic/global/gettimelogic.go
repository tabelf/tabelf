package global

import (
	"context"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"time"

	"github.com/zaihui/go-hutils"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTimeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTimeLogic {
	return &GetTimeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTimeLogic) GetTime(req *types.EmptyReq) (resp *types.TimeResp, err error) {
	return &types.TimeResp{Now: time.Now().Format(hutils.DateTimeLayout)}, nil
}
