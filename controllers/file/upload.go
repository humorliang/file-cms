package file

import (
	"github.com/gin-gonic/gin"
	"github.com/humorliang/file-cms/controllers/rsp"
)
//上传文件
func UplodeFile(c *gin.Context) {
	c.JSON(200, rsp.Success("上传成功"))
}
