package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
)

var ErrKey = errors.New("私钥不能为空")

func SignSHA256WithRSA(source string, privateKey *rsa.PrivateKey) (signature string, err error) {
	if privateKey == nil {
		return "", ErrKey
	}
	h := sha256.New()
	_, err = h.Write([]byte(source))
	if err != nil {
		return "", fmt.Errorf("签名失败: %w", err)
	}
	hashed := h.Sum(nil)
	signatureByte, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", fmt.Errorf("签名失败: %w", err)
	}
	return base64.StdEncoding.EncodeToString(signatureByte), nil
}
