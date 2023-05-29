package workspace

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"tabelf/backend/service/api/internal/logic/customer/workspace"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/api/pkg/utils"
)

func AddWorkspaceHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddWorkspaceRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := workspace.NewAddWorkspaceLogic(r.Context(), ctx)
		resp, err := l.AddWorkspace(&req)
		if err != nil {
			utils.HandleResponseError(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
