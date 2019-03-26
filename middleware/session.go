package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/humorliang/file-cms/controllers/rsp"
	"github.com/humorliang/file-cms/comm/e"
)

//检验session
func SessionAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user_id")
		if user == nil {
			c.JSON(400, rsp.Fails(e.SESSION_ERROR,e.GetMsg(e.SESSION_ERROR)))
			c.Abort()
		} else {
			c.Next()
		}
	}
}
