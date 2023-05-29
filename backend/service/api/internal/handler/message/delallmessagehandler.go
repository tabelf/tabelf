package message

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tabelf/backend/service/api/internal/logic/message"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"tabelf/backend/service/api/pkg/utils"
)

func DelAllMessageHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelAllMessageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := message.NewDelAllMessageLogic(r.Context(), ctx)
		resp, err := l.DelAllMessage(&req)
		if err != nil {
			utils.HandleResponseError(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
