package article

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tabelf/backend/service/api/internal/logic/article"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"tabelf/backend/service/api/pkg/utils"
)

func GetGoodArticleCollectionHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetGoodArticleCollectionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := article.NewGetGoodArticleCollectionLogic(r.Context(), ctx)
		resp, err := l.GetGoodArticleCollection(&req)
		if err != nil {
			utils.HandleResponseError(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
