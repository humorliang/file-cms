package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"fmt"
	"github.com/humorliang/file-cms/routers"
)

const (
	serverAddr     = "127.0.0.1:8080"
	readTimeout    = 10
	writeTimeout   = 20
	maxHeaderBytes = 1 << 20
)

func init() {
	fmt.Printf(`
	************** server is  run **************
	**** server addr : %s ***
	********************************************
	`, serverAddr)

}

func main() {
	gin.SetMode("release")
	//基础路由
	baseRouter := gin.New()
	routers.RegisterRouter(baseRouter)

	//自定义服务
	s := &http.Server{
		Addr:           serverAddr,
		Handler:        baseRouter,
		ReadTimeout:    readTimeout * time.Second,
		WriteTimeout:   writeTimeout * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}
	s.ListenAndServe()
}
