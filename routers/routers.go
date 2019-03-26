package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/humorliang/file-cms/controllers/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/humorliang/file-cms/controllers/file"
	"net/http"
)

func RegisterRouter(r *gin.Engine) {
	//session管理器
	store := cookie.NewStore([]byte("secretKey"))
	r.Use(sessions.Sessions("session_id", store))
	//异常恢复中间件
	r.Use(gin.Recovery())
	r.StaticFS("/static/", http.Dir("."))
	//API版本
	rV1 := r.Group("/v1")

	//无需认证
	rV1.POST("/user/login", user.Login)
	rV1.POST("/user/register", user.Register)

	//需要认证
	//rAuth := rV1.Group("/auth", middleware.SessionAuth())
	rAuth := rV1.Group("/auth")
	rAuth.POST("/file/upload", file.UplodeFile)
	rAuth.GET("/files/:page_num", file.GetFileList)
	rAuth.GET("/file/download",file.DownloadFile)
}
