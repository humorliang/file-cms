package utils

import (
	"crypto/sha256"
	"unsafe"
	"fmt"
)

func NewStrToSHA256(msg string) string {
	h := sha256.New()
	h.Write([]byte(msg))
	data := h.Sum(nil)
	//使用指针转化提高效率
	return fmt.Sprintf("%x",*(*string)(unsafe.Pointer(&data)))
}
