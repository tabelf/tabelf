package weblink

import (
	"io/ioutil"
	"net/http"
	"tabelf/backend/service/app"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tabelf/backend/service/api/internal/logic/customer/weblink"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"tabelf/backend/service/api/pkg/utils"
)

func AddPersonalLocalFileLinkHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddPersonalLocalFileLinkRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		if err := r.ParseMultipartForm(app.MaxLocalFileBufferUploadMemory); err != nil {
			httpx.Error(w, err)
			return
		}
		file, header, err := r.FormFile("file")
		if err != nil {
			httpx.Error(w, err)
			return
		}
		if header.Size > app.MaxLocalFileUploadMemoryLimit {
			utils.HandleResponseError(w, app.ErrCustomerLocalFileLimit(r.Context()))
			return
		}
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		req.File = bytes
		req.Filename = header.Filename
		req.FileType = header.Header.Get("Content-Type")

		l := weblink.NewAddPersonalLocalFileLinkLogic(r.Context(), ctx)
		resp, err := l.AddPersonalLocalFileLink(&req)
		if err != nil {
			utils.HandleResponseError(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
