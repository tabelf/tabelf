package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

var (
	ErrPrivateDecode = errors.New("解析私钥出错")
	ErrPubDecode     = errors.New("解析公钥出错")
	ErrPubType       = errors.New("不是公钥格式")
	ErrRSAType       = errors.New("不是RSA格式私钥")
)

// LoadPrivateKey 通过私钥的文本内容加载私钥.
func LoadPrivateKey(privateKeyStr string) (privateKey *rsa.PrivateKey, err error) {
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		return nil, ErrPrivateDecode
	}
	if block.Type != "RSA PRIVATE KEY" {
		return nil, ErrRSAType
	}
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("解析私钥出错: %w", err)
	}
	return key, nil
}

// LoadPublicKey 通过公钥的文本内容加载公钥.
func LoadPublicKey(publicKeyStr string) (publicKey *rsa.PublicKey, err error) {
	block, _ := pem.Decode([]byte(publicKeyStr))
	if block == nil {
		return nil, ErrPubDecode
	}
	if block.Type != "PUBLIC KEY" {
		return nil, ErrPubType
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("解析公钥出错: %w", err)
	}
	publicKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("%w: 不是RSA公钥", ErrPubDecode)
	}
	return publicKey, nil
}

// LoadPublicKeyWithoutBlock 通过公钥的文本内容加载公钥.
func LoadPublicKeyWithoutBlock(publicKeyStr string) (publicKey *rsa.PublicKey, err error) {
	value := fmt.Sprintf("-----BEGIN PUBLIC KEY-----\n%s\n-----END PUBLIC KEY-----", publicKeyStr)
	return LoadPublicKey(value)
}
