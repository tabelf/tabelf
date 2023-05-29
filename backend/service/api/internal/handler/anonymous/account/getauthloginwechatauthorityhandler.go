package account

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"tabelf/backend/service/api/internal/logic/anonymous/account"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/api/pkg/utils"
)

func GetAuthLoginWechatAuthorityHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAuthLoginWechatAuthorityRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := account.NewGetAuthLoginWechatAuthorityLogic(r.Context(), ctx)
		resp, err := l.GetAuthLoginWechatAuthority(&req)
		if err != nil {
			utils.HandleResponseError(w, err)
		} else {
			if _, err = w.Write([]byte(resp.Echostr)); err != nil {
				return
			}
		}
	}
}
