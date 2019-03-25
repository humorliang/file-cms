package file

import (
	"github.com/gin-gonic/gin"
	"github.com/humorliang/file-cms/controllers/rsp"
	"github.com/gin-contrib/sessions"
	"fmt"
)
//上传文件
func UplodeFile(c *gin.Context) {
	session:=sessions.Default(c)
	fmt.Println(session.Get("user_id"))
	c.JSON(200, rsp.Success("上传成功"))
}
