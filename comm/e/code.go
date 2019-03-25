package e

const (
	INTERSERVER_ERROR = 00000
	PARAM_PARSE       = 00001

	//用户错误信息 10xxx
	USER_NOT_EXIST  = 10001
	USER_EXIST      = 10002
	USER_LOGIN_FAIL = 10003
	USER_REGISTER_FAIL=10004
)
/*
var msg = map[int]string{
	INTERSERVER_ERROR: "服务器异常",
	PARAM_PARSE:       "请求参数错误",
	USER_NOT_EXIST:    "用户不存在",
	USER_EXIST:        "用户已存在",
	USER_LOGIN_FAIL:"用户名或密码错误",
}
*/

