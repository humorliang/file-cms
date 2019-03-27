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
	"crypto/aes"
	"crypto/cipher"
	"bytes"
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

/*
AES:
全称Cipher Block Chaining mode，密码分组链接模式。
分组方式：将明文分组与前一个密文分组进行XOR运算，然后再进行加密。每个分组的加解密都依赖于前一个分组。
而第一个分组没有前一个分组，因此需要一个初始化向量（initialization vector）。
CBC模式+PKCS5 填充方式 实现AES加密
*/

//AES key必须是16字节 24字节 32字节中的一种
func CheckAesKey(strKey string) []byte {
	keyLen := len(strKey)
	arrKey := []byte(strKey)
	if keyLen >= 32 {
		//取前32个字节
		return arrKey[:32]
	}
	if keyLen >= 24 {
		//取前24个字节
		return arrKey[:24]
	}
	if keyLen >= 16 {
		//取前16个字节
		return arrKey[:16]
	}
	//补齐16个字节
	tmp := make([]byte, 16)
	for i := 0; i < 16; i++ {
		if i < keyLen {
			tmp[i] = arrKey[i]
		} else {
			tmp[i] = '0'
		}
	}
	return tmp
}

//AES加密的默认向量，
const aesIvDefValue = "217c1b45372d45c3"

//AES加密
func AesEncrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("err=", err)
		return nil, errors.New("invalid decrypt key")
	}
	blockSize := block.BlockSize()
	plaintext = PKCS5Padding(plaintext, blockSize)
	iv := []byte(aesIvDefValue)
	blockMode := cipher.NewCBCEncrypter(block, iv)

	ciphertext := make([]byte, len(plaintext))
	blockMode.CryptBlocks(ciphertext, plaintext)

	return ciphertext, nil
}

//AES解密
func AesDecrypt(ciphertext []byte, key []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("invalid decrypt key")
	}

	blockSize := block.BlockSize()

	if len(ciphertext) < blockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := []byte(aesIvDefValue)
	if len(ciphertext)%blockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	blockModel := cipher.NewCBCDecrypter(block, iv)

	plaintext := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plaintext, ciphertext)
	plaintext = PKCS5UnPadding(plaintext)

	return plaintext, nil
}

//组装
func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

//拆分
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
