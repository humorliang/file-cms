package file

import (
	"github.com/gin-gonic/gin"
	"github.com/humorliang/file-cms/controllers/rsp"
	"github.com/humorliang/file-cms/comm/e"
	"fmt"
	"github.com/humorliang/file-cms/db"
	"github.com/humorliang/file-cms/utils"
	"github.com/humorliang/file-cms/comm/g"
)

//上传文件
func UplodeFile(c *gin.Context) {
	//上传文件
	file, name, err := c.Request.FormFile("file")
	if name == nil {
		c.JSON(200, rsp.Fails(e.PARAM_PARSE, e.GetMsg(e.PARAM_PARSE)))
		return
	}
	if err != nil {
		fmt.Println(err)
		c.JSON(500, rsp.Fails(e.INTERSERVER_ERROR, e.GetMsg(e.INTERSERVER_ERROR)))
		return
	}
	//设置缓冲区
	buf := make([]byte, name.Size)
	_, err = file.Read(buf)
	if err != nil {
		c.JSON(500, rsp.Fails(e.INTERSERVER_ERROR, e.GetMsg(e.INTERSERVER_ERROR)))
		return
	}
	bufEncrypt, err := utils.RsaEncrypt(buf, g.RsaKey)
	res, err := db.AddFile(name.Filename, bufEncrypt)
	if err != nil {
		c.JSON(500, rsp.Fails(e.INTERSERVER_ERROR, e.GetMsg(e.INTERSERVER_ERROR)))
		return
	} else {
		c.JSON(201, rsp.Success(gin.H{
			"file_id":   res.FileId,
			"file_name": name.Filename,
			"file_hash": utils.SHA256Hash(buf),
		}))
	}
}
