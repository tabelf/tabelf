package middleware

import "net/http"

type CorsMiddleware struct{}

func NewCorsMiddleware() *CorsMiddleware {
	return &CorsMiddleware{}
}

func (m *CorsMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setHeader(w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func (m *CorsMiddleware) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setHeader(w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers",
		"Content-Type, X-CSRF-Token, Authorization, AccessToken, Token")
	w.Header().Set("Access-Control-Allow-Methods",
		"GET, POST, PUT, DELETE, OPTIONS, PATCH")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Length, "+
		"Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}
