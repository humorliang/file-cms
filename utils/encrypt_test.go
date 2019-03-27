package utils

import (
	"testing"
	"fmt"
)

func TestNewStrToSHA256(t *testing.T) {
	fmt.Printf("%x", NewStrToSHA256("test"))
	t.Log(NewStrToSHA256("test") == NewStrToSHA256("test"))
}

func TestAesEncryptAndAesDecrypt(t *testing.T) {
	data := "this is data"
	key := CheckAesKey("123123kj131k3k14")
	encryptData, err := AesEncrypt([]byte(data), key)
	fmt.Println("aes加密：", encryptData, err)
	decryptData, err := AesDecrypt(encryptData, key)
	fmt.Println("aes解密：", string(decryptData[:]), err)
}
