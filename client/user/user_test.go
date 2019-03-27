package user

import (
	"testing"
	"fmt"
)

func TestReqLogin(t *testing.T) {
	code, data, err := ReqLogin("user", "123456")
	fmt.Println(code, string(data[:]), err)
}

func TestReqRegister(t *testing.T) {
	code, data, err := ReqRegister("user2", "123456")
	fmt.Println(code, string(data[:]), err)
}
