package file

import (
	"github.com/gin-gonic/gin"
	"net/url"
	"github.com/humorliang/file-cms/controllers/rsp"
	"github.com/humorliang/file-cms/comm/e"
	"github.com/humorliang/file-cms/db"
	"unsafe"
	"fmt"
	"strconv"
	"github.com/humorliang/file-cms/utils"
	"github.com/humorliang/file-cms/comm/g"
)

//文件下载
func DownloadFile(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(200, rsp.Fails(e.PARAM_PARSE, e.GetMsg(e.PARAM_PARSE)))
		return
	}
	file, err := db.GetFile(*(*int64)(unsafe.Pointer(&id)))
	if err != nil {
		fmt.Println(err)
		c.JSON(500, rsp.Fails(e.INTERSERVER_ERROR, e.GetMsg(e.INTERSERVER_ERROR)))
		return
	}
	fData, err := utils.RsaEncrypt(file.FileData, g.RsaKey)
	c.Header("Content-Disposition", "attachment; filename="+url.QueryEscape(file.FileName))
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Expires", "0")
	c.Header("Cache-Control", "must-revalidate")
	c.Header("File-Hash", utils.SHA256Hash(fData))
	c.Header("Pragma", "public")
	c.Data(200, "application/octet-stream", fData)
}
