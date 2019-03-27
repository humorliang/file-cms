package utils

import (
	"testing"
	"fmt"
)

func TestCheckExist(t *testing.T) {
	b := CheckPathExist("C:\\Users\\jackson\\go\\src\\github.com\\humorliang\\file-cms\\client\\test\\a.docx")
	fmt.Println(b)
}
