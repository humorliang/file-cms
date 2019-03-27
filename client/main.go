package main

import (
	"flag"
	"github.com/humorliang/file-cms/client/base"
	"fmt"
	"github.com/humorliang/file-cms/client/user"
	"os"
	"github.com/humorliang/file-cms/client/file"
)

func init() {
	helpInfo := `
	**************************************************
	******* 文件管理系统手册 **********
	******* 1. 用户登陆	-login     "username,password" 
	******* 2. 用户注册	-register  "username,password"
	******* 3. 上传文件	-upload    "fileAbsPath"	
	******* 4. 文件列表	-files     "pageNum"
	******* 5. 文件下载	-download  "文件ID"
	**************************************************
`
	fmt.Println(helpInfo)
}
func main() {

	//登陆
	var loginInfo base.ArrayFlags
	flag.Var(&loginInfo, "login", "this is login")

	//注册
	var registerInfo base.ArrayFlags
	flag.Var(&registerInfo, "register", "this is register")

	//上传
	uploadFilePath := flag.String("upload", "", "this is upload")

	//列表
	files := flag.String("files", "1", "this is get file list")

	//下载
	download := flag.String("download", "1", "this is download file")

	//参数解析
	flag.Parse()

	//登陆
	if loginInfo != nil {
		code, data, err := user.ReqLogin(loginInfo[0], loginInfo[1])
		if err != nil {
			fmt.Println("login error:", err)
			os.Exit(1)
		}
		fmt.Println("登陆成功：\n", code, data)
	}

	//注册
	if registerInfo != nil {
		code, data, err := user.ReqRegister(registerInfo[0], loginInfo[1])
		if err != nil {
			fmt.Println("register error:", err)
			os.Exit(1)
		}
		fmt.Println("注册成功：\n", code, data)
	}

	//上传
	if *uploadFilePath != "" {
		code, data, err := file.UploadFile(*uploadFilePath)
		if err != nil {
			fmt.Println("upload file error:", err)
			os.Exit(1)
		}
		fmt.Println("上传成功：\n", code, data)
	}

	//下载
	if *download != "" {
		code, data, err := file.DownLoadFile(*uploadFilePath)
		if err != nil {
			fmt.Println("download file error:", err)
			os.Exit(1)
		}
		fmt.Println("下载成功：\n", code, data)
	}

	//文件列表
	if *files != "" {
		code, data, err := file.FileList(*uploadFilePath)
		if err != nil {
			fmt.Println("show file list error:", err)
			os.Exit(1)
		}
		fmt.Println("获取成功：\n", code, data)
	}
}
