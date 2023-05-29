package folder

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tabelf/backend/service/api/internal/logic/customer/folder"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"tabelf/backend/service/api/pkg/utils"
)

func CopyPersonalFolderHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CopyPersonalFolderRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := folder.NewCopyPersonalFolderLogic(r.Context(), ctx)
		resp, err := l.CopyPersonalFolder(&req)
		if err != nil {
			utils.HandleResponseError(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
