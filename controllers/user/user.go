package user

import (
	"github.com/gin-gonic/gin"
	"github.com/humorliang/file-cms/controllers/rsp"
	"github.com/humorliang/file-cms/comm/e"
	"github.com/humorliang/file-cms/db"
	"unsafe"
	"github.com/gin-contrib/sessions"
)

//登陆
func Login(c *gin.Context) {
	uName := c.PostForm("user_name")
	pwd := c.PostForm("pass_word")
	if uName == "" || pwd == "" {
		c.JSON(200, rsp.Fails(e.PARAM_PARSE, e.GetMsg(e.PARAM_PARSE)))
		return
	} else {
		if u, err := db.GetUserByName(uName); u == nil {
			if err != nil {
				c.JSON(500, rsp.Fails(e.INTERSERVER_ERROR, e.GetMsg(e.INTERSERVER_ERROR)))
				return
			} else {
				c.JSON(200, rsp.Fails(e.USER_LOGIN_FAIL, e.GetMsg(e.USER_LOGIN_FAIL)))
				return
			}
		} else {
			session := sessions.Default(c)
			session.Set("user_id", *(*int64)(unsafe.Pointer(&u.UserId)))
			session.Save()
			c.JSON(200, rsp.Success(gin.H{
				"user_id": *(*int64)(unsafe.Pointer(&u.UserId)),
			}))
		}
	}
}

//用户注册
func Register(c *gin.Context) {
	//表单输入
	uName := c.PostForm("user_name")
	pwd := c.PostForm("pass_word")
	if uName == "" || pwd == "" {
		c.JSON(200, rsp.Fails(e.PARAM_PARSE, e.GetMsg(e.PARAM_PARSE)))
		return
	} else {
		if u, err := db.GetUserByName(uName); u != nil {
			if err != nil {
				c.JSON(500, rsp.Fails(e.USER_REGISTER_FAIL, e.GetMsg(e.USER_REGISTER_FAIL)))
				return
			} else {
				c.JSON(200, rsp.Fails(e.USER_EXIST, e.GetMsg(e.USER_EXIST)))
				return
			}
		} else {
			user := &db.User{UserName: uName, PassWord: pwd}
			u, err := db.AddUser(user)
			if err != nil {
				c.JSON(500, rsp.Fails(e.USER_REGISTER_FAIL, e.GetMsg(e.USER_REGISTER_FAIL)))
			} else {
				c.JSON(201, rsp.Success(gin.H{
					"user_id": *(*int64)(unsafe.Pointer(&u.UserId)),
				}))
			}
		}
	}
}
