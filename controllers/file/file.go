package file

import (
	"github.com/gin-gonic/gin"
	"github.com/humorliang/file-cms/db"
	"github.com/humorliang/file-cms/controllers/rsp"
	"github.com/humorliang/file-cms/comm/e"
)

//获取文件列表
func GetFileList(c *gin.Context) {
	files, err := db.GetFiles()
	if err != nil {
		c.JSON(500, rsp.Fails(e.INTERSERVER_ERROR, e.GetMsg(e.INTERSERVER_ERROR)))
		return
	}
	rspInfo := []gin.H{}
	for _, v := range files.FList {
		rspInfo = append(rspInfo, gin.H{
			"file_id":     v.FileId,
			"file_name":   v.FileName,
			"create_date": v.CreateDate,
		})
	}
	c.JSON(200, rspInfo)
}
