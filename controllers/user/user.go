package user

import (
	"github.com/gin-gonic/gin"
	"github.com/humorliang/file-cms/controllers/rsp"
	"github.com/gin-contrib/sessions"
)

//登陆
func Login(c *gin.Context) {
	//获取会话
	session := sessions.Default(c)

	session.Set("user_id", "1")
	session.Save()
	c.JSON(200, rsp.Success("登陆成功"))
}

//用户注册
func Register(c *gin.Context) {
	c.JSON(200, rsp.Success("注册成功"))
}
