package middleware

import (
	"context"
	"errors"
	"net/http"

	"tabelf/backend/service/api/pkg/utils"
	"tabelf/backend/service/extensions"
)

type JWTInfo struct {
	UID string
}

var ErrNotHasToken = errors.New("not has token")

type JWTKey string

type JWTMiddleware struct {
	extensions.JWTConf
}

func NewJWTMiddleware(conf extensions.JWTConf) *JWTMiddleware {
	return &JWTMiddleware{conf}
}

func NewJWTKey() JWTKey {
	return "user"
}

func (m *JWTMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "未提供有效身份许可", http.StatusUnauthorized)
			return
		}
		tokenType, claims := utils.IsValidateOfToken(token, m.JwtKey)
		var resultMessage string
		reqCtx := r.Context()
		switch tokenType {
		case utils.ValidateToken:
			{
				r = r.WithContext(context.WithValue(reqCtx, NewJWTKey(), JWTInfo{
					UID: claims.UID,
				}))
				next(w, r)
				return
			}
		case utils.ExpiredToken:
			resultMessage = "令牌已失效!"
		case utils.BadToken:
			resultMessage = "未提供有效令牌"
		}
		http.Error(w, resultMessage, http.StatusUnauthorized)
	}
}

func ContextElement(ctx context.Context) (*JWTInfo, error) {
	if info, ok := ctx.Value(NewJWTKey()).(JWTInfo); ok {
		return &info, nil
	}
	return nil, ErrNotHasToken
}
