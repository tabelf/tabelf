package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// AesDecrypt AES解密.
func AesDecrypt(encryptedData, initialVector, sessionKey string) ([]byte, error) {
	var sessionKeyByte []byte
	var encryptedDataByte []byte
	var iv []byte
	var err error
	sessionKeyByte, err = base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}
	iv, err = base64.StdEncoding.DecodeString(initialVector)
	if err != nil {
		return nil, err
	}
	encryptedDataByte, err = base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(sessionKeyByte)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(encryptedDataByte))
	blockMode.CryptBlocks(origData, encryptedDataByte)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
