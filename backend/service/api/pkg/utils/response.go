package utils

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tabelf/backend/common"
	"tabelf/backend/service/api/internal/types"
)

func HandleResponseError(w http.ResponseWriter, err error) {
	respErr := &types.ErrorMessage{
		Code: http.StatusBadRequest,
	}
	var zErr *common.ZError
	if ok := errors.As(err, &zErr); ok {
		if code, err := strconv.Atoi(zErr.Code); err == nil {
			respErr.Code = code
		}
		respErr.TraceID = zErr.TraceID
		respErr.SpanID = zErr.SpanID
		respErr.Message = zErr.Message
	} else {
		respErr.Message = err.Error()
	}
	httpx.WriteJson(w, http.StatusBadRequest, respErr)
}
