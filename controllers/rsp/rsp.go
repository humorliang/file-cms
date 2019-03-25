package rsp

import (
	"github.com/gin-gonic/gin"
	"github.com/humorliang/file-cms/comm/e"
)

//响应成功
func Success(data interface{}) gin.H {
	return gin.H{
		"code": 0,
		"msg":  "SUCCESS",
		"data": data,
	}
}

//响应失败
func Fails(eCode int, data interface{}) gin.H {
	return gin.H{
		"code": eCode,
		"msg":  e.GetMsg(eCode),
		"data": data,
	}
}
