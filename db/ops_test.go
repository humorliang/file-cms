package db

import (
	"testing"
	"github.com/humorliang/file-cms/comm/g"
	"fmt"
)

func TestGetUserByName(t *testing.T) {
	g.InitDB()
	//user,err :=GetUserByName("user")
	//fmt.Println(user,err)

}
func TestAddUser(t *testing.T) {
	g.InitDB()
	user := &User{UserName: "user", PassWord: "123"}
	u, err := AddUser(user)
	fmt.Println(u, err)
}

func TestGetFile(t *testing.T) {
	g.InitDB()
	f, err := GetFile(3)
	fmt.Println(f, err)
}

func TestGetFiles(t *testing.T) {
	g.InitDB()
	f, err := GetFiles()
	fmt.Println(f, err)
}
