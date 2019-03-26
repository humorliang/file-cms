package utils

import (
	"crypto/sha256"
	"unsafe"
	"fmt"
	"encoding/pem"
	"errors"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
)

func NewStrToSHA256(msg string) string {
	h := sha256.New()
	h.Write([]byte(msg))
	data := h.Sum(nil)
	//使用指针转化提高效率
	return fmt.Sprintf("%x", *(*string)(unsafe.Pointer(&data)))
}

func SHA256Hash(data []byte) string {
	h := sha256.New()
	h.Write(data)
	hData := h.Sum(nil)
	//使用指针转化提高效率
	return fmt.Sprintf("%x", *(*string)(unsafe.Pointer(&hData)))
}

//数据加密
func RsaEncrypt(origData []byte, publicKey string) ([]byte, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

//数据解密
func RsaDecrypt(ciphertext []byte, privateKey string) ([]byte, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
