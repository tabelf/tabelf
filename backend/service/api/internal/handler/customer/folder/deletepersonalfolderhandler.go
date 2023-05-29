package folder

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"tabelf/backend/service/api/internal/logic/customer/folder"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/api/pkg/utils"
)

func DeletePersonalFolderHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeletePersonalFolderRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := folder.NewDeletePersonalFolderLogic(r.Context(), ctx)
		resp, err := l.DeletePersonalFolder(&req)
		if err != nil {
			utils.HandleResponseError(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
