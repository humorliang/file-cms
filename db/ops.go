package db

import (
	"github.com/humorliang/file-cms/comm/g"
	"database/sql"
	"github.com/humorliang/file-cms/utils"
	"fmt"
)

type User struct {
	UserId     int64  `json:"user_id"`
	UserName   string `json:"user_name"`
	PassWord   string `json:"pass_word,omitempty"`
	CreateDate string `json:"create_date"`
}

type File struct {
	FileId     int64  `json:"file_id"`
	FileName   string `json:"file_name"`
	FileData   []byte `json:"file_data"`
	CreateDate string `json:"create_date"`
}
type Files struct {
	FList []File
}

//获取用户信息通过用户名
func GetUserByName(name string) (*User, error) {
	user := new(User)
	//fmt.Println(user)
	sqlStatement := `SELECT user_id FROM users WHERE user_name=$1`
	row := g.PsgCon.QueryRow(sqlStatement, name)
	//fmt.Println(row)
	err := row.Scan(&user.UserId)

	//fmt.Println("get name ", user, err)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return user, nil
	default:
		return nil, err
	}
}

//添加用户信息
func AddUser(user *User) (*User, error) {
	//postgres没有返回ID可以使用RETURNING
	sqlStatement := `INSERT INTO users (user_name,pass_word) VALUES ($1,$2) RETURNING user_id,create_date`
	err := g.PsgCon.QueryRow(sqlStatement, user.UserName, utils.NewStrToSHA256(user.PassWord)).Scan(&user.UserId,
		&user.CreateDate)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return user, err
}

//验证用户
func GetUserByNameAndPwd(name string, pwd string) (*User, error) {
	user := new(User)
	sqlStatement := `SELECT user_id FROM users WHERE user_name=$1 AND pass_word=$2`
	row := g.PsgCon.QueryRow(sqlStatement, name, utils.NewStrToSHA256(pwd))
	err := row.Scan(&user.UserId)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return user, nil
	default:
		return nil, err
	}
}

//文件写入数据库
func AddFile(name string, data []byte) (*File, error) {
	sqlStatement := `INSERT INTO files(file_name,file_data) VALUES ($1,$2) RETURNING file_id`
	file := File{}
	err := g.PsgCon.QueryRow(sqlStatement, name, data).Scan(&file.FileId)
	if err != nil {
		return nil, err
	}
	return &file, nil
}

//读取文件数据
func GetFile(id int64) (*File, error) {
	fmt.Println(id)
	sqlStatement := `SELECT file_name,file_data FROM files WHERE file_id=$1`
	file := File{}
	err := g.PsgCon.QueryRow(sqlStatement, id).Scan(&file.FileName, &file.FileData)
	if err != nil {
		return nil, err
	}
	file.FileId = id
	return &file, nil
}

//获取文件列表
func GetFiles() (*Files, error) {
	sqlStatement := `SELECT file_id,file_name,file_data,create_date FROM files`
	rows, err := g.PsgCon.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	files := Files{}
	for rows.Next() {
		var f File
		err = rows.Scan(&f.FileId, &f.FileName, &f.FileData, &f.CreateDate)
		if err != nil {
			return nil, err
		}
		files.FList = append(files.FList, f)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &files, nil
}
