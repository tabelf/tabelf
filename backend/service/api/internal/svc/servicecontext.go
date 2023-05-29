package svc

import (
	"net/http"
	"tabelf/backend/service/api/internal/middleware"
	"tabelf/backend/service/api/pkg/swagger"
	"tabelf/backend/service/app"

	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config  app.Config
	Swagger http.HandlerFunc
	JWT     rest.Middleware
}

func NewServiceContext(c app.Config, env string) *ServiceContext {
	sct := &ServiceContext{
		Config:  c,
		Swagger: swagger.Doc("/swagger", env),
		JWT:     middleware.NewJWTMiddleware(c.Jwt).Handle,
	}
	return sct
}
