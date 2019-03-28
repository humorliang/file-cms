package file

import (
	"github.com/gin-gonic/gin"
	"github.com/humorliang/file-cms/db"
	"github.com/humorliang/file-cms/controllers/rsp"
	"github.com/humorliang/file-cms/comm/e"
	"strconv"
	"github.com/humorliang/file-cms/comm/logging"
)

//获取文件列表
func GetFileList(c *gin.Context) {
	pageNum := c.Param("page_num")
	num, err := strconv.Atoi(pageNum)
	if err != nil {
		c.JSON(400, rsp.Fails(e.PARAM_PARSE, e.GetMsg(e.PARAM_PARSE)))
		logging.Error(err)
		return
	}
	files, err := db.GetFiles(num)
	if err != nil {
		c.JSON(500, rsp.Fails(e.INTERSERVER_ERROR, e.GetMsg(e.INTERSERVER_ERROR)))
		logging.Error(err)
		return
	}
	if len(files.FList) == 0 {
		c.JSON(200, rsp.Success("无更多数据"))
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
	c.JSON(200, rsp.Success(rspInfo))
}
