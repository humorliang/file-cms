package e

const (
	INTERSERVER_ERROR = 50000
	PARAM_PARSE       = 50001

	//用户错误信息 10xxx
	USER_NOT_EXIST     = 10001
	USER_EXIST         = 10002
	USER_LOGIN_FAIL    = 10003
	USER_REGISTER_FAIL = 10004

	//验证类错误 40xxx
	SESSION_ERROR = 40001
)

var msg = map[int]string{
	INTERSERVER_ERROR:  "服务器异常",
	PARAM_PARSE:        "请求参数错误",
	USER_NOT_EXIST:     "用户不存在",
	USER_EXIST:         "用户已存在",
	USER_LOGIN_FAIL:    "用户名或密码错误",
	USER_REGISTER_FAIL: "用户注册失败",
	//
	SESSION_ERROR: "Session无效",
}
