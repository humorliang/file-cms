package utils

import (
	"testing"
	"fmt"
)

func TestGenRsaKey(t *testing.T) {
	err := GenRsaKey(1024)
	fmt.Println(err)
}
