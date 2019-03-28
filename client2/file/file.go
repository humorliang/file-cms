package file

import (
	"github.com/humorliang/file-cms/utils"
	"errors"
	"bytes"
	"mime/multipart"
	"os"
	"io"
	"net/http"
	"io/ioutil"
	"path/filepath"
	"strings"
)

//上传
func UploadFile(filePath string) (code int, data []byte, err error) {
	//检查文件
	if !utils.CheckPathExist(filePath) {
		return 0, nil, errors.New("文件不存在")
	}
	//获取文件名
	_, fileName := filepath.Split(filePath)
	//创建文件表单
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	formFile, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return 0, nil, err
	}
	//读取文件数据
	srcFile, err := os.Open(filePath)
	if err != nil {
		return 0, nil, err
	}
	defer srcFile.Close()
	//拷贝数据
	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		return 0, nil, err
	}
	//发送数据
	contentType := writer.FormDataContentType()
	//关闭结尾行
	writer.Close()
	//发送请求
	rsp, err := http.Post("http://127.0.0.1:8080/v1/auth/file/upload", contentType, buf)
	if err != nil {
		return 0, nil, err
	}
	defer rsp.Body.Close()
	data, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return 0, nil, err
	}
	//读取数据
	return rsp.StatusCode, data, nil
}

//下载
func DownLoadFile(id string) (code int, data []byte, err error) {
	rsp, err := http.Get("http://127.0.0.1:8080/v1/auth/file/download?id=" + id)
	defer rsp.Body.Close()

	if rsp.StatusCode != 200 {
		data, err = ioutil.ReadAll(rsp.Body)
		if err != nil {
			return rsp.StatusCode, nil, err
		}
		return rsp.StatusCode, data, err
	} else {
		fileDesp := rsp.Header.Get("Content-Disposition")
		if fileDesp == "" {
			return
		}
		fileName := strings.Split(fileDesp, "=")[1:][0]
		file, err := os.Create(fileName)
		_, err = io.Copy(file, rsp.Body)
		if err != nil {
			return 0, nil, err
		}
		return rsp.StatusCode, []byte("下载成功"), nil
	}

}

//文件列表
func FileList(pageNum string) (code int, data []byte, err error) {
	rsp, err := http.Get("http://127.0.0.1:8080/v1/auth/files/" + pageNum)
	defer rsp.Body.Close()
	data, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return 0, data, err
	}
	return rsp.StatusCode, data, err
}
