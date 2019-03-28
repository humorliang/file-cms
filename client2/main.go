package main

import (
	"github.com/humorliang/file-cms/client2/cmd"
	"fmt"
	"os"
	"github.com/humorliang/file-cms/client2/user"
	"github.com/humorliang/file-cms/client2/file"
)

func main() {
	cmd.Execute()
	//登陆
	if cmd.LoginSwitch {
		code, data, err := user.ReqLogin(cmd.LgUserName, cmd.LgPassWord)
		if err != nil {
			fmt.Println("login error:", err)
			os.Exit(1)
		}
		fmt.Println("登陆成功：\n", code, string(data[:]))
	}

	//注册
	if cmd.RegisterSwitch {

		code, data, err := user.ReqRegister(cmd.RgUserName, cmd.RgPassWord)
		if err != nil {
			fmt.Println("register error:", err)
			os.Exit(1)
		}
		fmt.Println("注册信息：\n", code, string(data[:]))
	}

	//上传
	if cmd.UploadSwitch {
		code, data, err := file.UploadFile(cmd.FileAbsPath)
		if err != nil {
			fmt.Println("upload file error:", err)
			os.Exit(1)
		}
		fmt.Println("上传成功：\n", code, string(data[:]))
	}

	//下载
	if cmd.DownloadSwitch {
		code, data, err := file.DownLoadFile(cmd.FileId)
		if err != nil {
			fmt.Println("download file error:", err)
			os.Exit(1)
		}
		fmt.Println("下载成功：\n", code, string(data[:]))
	}

	//文件列表
	if cmd.FilesSwitch {
		code, data, err := file.FileList(cmd.PageNum)
		if err != nil {
			fmt.Println("show file list error:", err)
			os.Exit(1)
		}
		fmt.Println("获取成功：\n", code, string(data[:]))
	}
}
