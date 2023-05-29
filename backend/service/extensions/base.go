package extensions

import (
	"net/http"
)

type Extension interface {
	Init() error
	Close() error
	OpenAPIMiddleWare(next http.Handler) http.Handler
}
