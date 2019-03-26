package utils

import (
	"testing"
	"fmt"
)

func TestNewStrToSHA256(t *testing.T) {
	fmt.Printf("%x",NewStrToSHA256("test"))
	t.Log(NewStrToSHA256("test")==NewStrToSHA256("test"))
}