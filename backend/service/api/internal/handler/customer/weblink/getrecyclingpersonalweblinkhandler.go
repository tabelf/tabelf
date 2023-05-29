package weblink

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"tabelf/backend/service/api/internal/logic/customer/weblink"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/api/pkg/utils"
)

func GetRecyclingPersonalWebLinkHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRecyclingPersonalWebLinkRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := weblink.NewGetRecyclingPersonalWebLinkLogic(r.Context(), ctx)
		resp, err := l.GetRecyclingPersonalWebLink(&req)
		if err != nil {
			utils.HandleResponseError(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
