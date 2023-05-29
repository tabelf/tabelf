package utils

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenType int

// token的状态.
const (
	ValidateToken TokenType = 0
	BadToken      TokenType = 1
	ExpiredToken  TokenType = 2
)

const AudienceUser = "user"

type JWTClaims struct {
	UID string `json:"uid"`
	jwt.StandardClaims
}

func GetJwtToken(secretKey, issuer string, seconds int64, claims JWTClaims) (string, error) {
	expireTime := time.Now().Add(time.Duration(seconds) * time.Second)
	claims.ExpiresAt = expireTime.Unix()
	claims.Issuer = issuer

	aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := aToken.SignedString([]byte(secretKey))

	return token, err
}

func ParseToken(token string, secret string) (*JWTClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		strings.TrimPrefix(token, "Bearer "), &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*JWTClaims); ok {
			return claims, nil
		}
	}
	return nil, err
}

func IsValidateOfToken(token string, secret string) (TokenType, *JWTClaims) {
	claims, err := ParseToken(token, secret)
	if err != nil {
		return BadToken, nil
	}
	if claims.ExpiresAt-time.Now().Unix() < 0 {
		return ExpiredToken, nil
	}

	if err := claims.Valid(); err != nil {
		return BadToken, nil
	}

	return ValidateToken, claims
}
