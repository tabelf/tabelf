package swagger

import (
	"net/http"

	"tabelf/backend/service/api/internal/svc"
)

func SwagHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return ctx.Swagger
}
