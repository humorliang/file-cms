package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/humorliang/file-cms/comm/g"
	"github.com/humorliang/file-cms/comm/setting"
	"github.com/humorliang/file-cms/routers"
	"net/http"
	"time"
)

const (
	maxHeaderBytes = 1 << 20
)

func init() {
	fmt.Printf(`
	****   server is  running     **************
	**** server addr : %s:%s **********
	********************************************
	`, setting.ServerCfg.HttpAddr, setting.ServerCfg.HttpPort)
	//初始化操作
	g.InitDB()
	//设置开发模式
	gin.SetMode(setting.ServerCfg.RunMode)
}

func main() {
	//gin.SetMode("release")
	//基础路由
	baseRouter := gin.New()
	routers.RegisterRouter(baseRouter)

	//自定义服务
	s := &http.Server{
		Addr:           setting.ServerCfg.HttpAddr + ":" + setting.ServerCfg.HttpPort,
		Handler:        baseRouter,
		ReadTimeout:    setting.ServerCfg.ReadTimeout * time.Second,
		WriteTimeout:   setting.ServerCfg.WriteTimeout * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}
	s.ListenAndServe()
}
