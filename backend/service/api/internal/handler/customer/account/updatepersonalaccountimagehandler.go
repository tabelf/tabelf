package account

import (
	"io/ioutil"
	"net/http"
	"tabelf/backend/service/app"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tabelf/backend/service/api/internal/logic/customer/account"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"tabelf/backend/service/api/pkg/utils"
)

func UpdatePersonalAccountImageHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdatePersonalAccountImageRequest
		if err := httpx.ParsePath(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		if err := r.ParseMultipartForm(app.MaxBufferUploadMemory); err != nil {
			httpx.Error(w, err)
			return
		}
		file, header, err := r.FormFile("file")
		if err != nil {
			httpx.Error(w, err)
			return
		}
		if !app.StringSliceContains([]string{
			"image/jpeg",
			"image/png",
			"image/webp",
			"image/bmp",
			"image/gif",
		}, header.Header.Get("Content-Type")) {
			utils.HandleResponseError(w, app.ErrCustomerHeaderImageTypeInvalid(r.Context()))
			return
		}
		if header.Size > app.MaxUploadHeaderImageLimit {
			utils.HandleResponseError(w, app.ErrCustomerHeaderImageLimit(r.Context()))
			return
		}
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		req.File = bytes
		req.Filename = header.Filename

		l := account.NewUpdatePersonalAccountImageLogic(r.Context(), ctx)
		resp, err := l.UpdatePersonalAccountImage(&req)
		if err != nil {
			utils.HandleResponseError(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
