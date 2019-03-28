package file

import (
	"testing"
	"fmt"
)

func TestUploadFile(t *testing.T) {
	code, data, err := UploadFile(`C:\Users\jackson\go\src\github.com\humorliang\file-cms\client\test\test.docx`)
	fmt.Println(code, string(data[:]), err)
}

func TestDownLoadFile(t *testing.T) {
	code, data, err := DownLoadFile("22")
	fmt.Println(code, string(data[:]), err)
}

func TestFileList(t *testing.T) {
	code, data, err := FileList("1")
	fmt.Println(code, string(data[:]), err)
}
