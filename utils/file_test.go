package utils

import (
	"testing"
	"fmt"
	"os"
)

func TestCheckExist(t *testing.T) {
	b := CheckPathExist("C:\\Users\\jackson\\go\\src\\github.com\\humorliang\\file-cms\\client\\test\\a.docx")
	fmt.Println(b)
}

func TestMustOpenSrc(t *testing.T) {
	//获取runtime运行的文件夹
	//.\go\src\github.com\humorliang\file-cms\utils
	dir, err := os.Getwd()
	fmt.Println(dir, err)
}
