package global

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zaihui/go-hutils"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrBadRequest = errors.New("bad request")

type GetErrorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetErrorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetErrorLogic {
	return &GetErrorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetErrorLogic) GetError(req *types.ErrorReq) (resp *types.ErrorResp, err error) {
	if req.Code == http.StatusInternalServerError {
		panic("this is 500")
	}
	err = fmt.Errorf("%w: this is 400", ErrBadRequest)
	return nil, hutils.NewZError(l.ctx, http.StatusBadRequest, err.Error())
}
