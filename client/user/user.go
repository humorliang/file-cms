package user

import (
	"net/http"
	"net/url"
	"io/ioutil"
)

//发起登陆请求
func ReqLogin(userName string, passWord string) (code int, data []byte, err error) {
	rsp, err := http.PostForm("http://127.0.0.1:8080/v1/user/login",
		url.Values{"user_name": {userName}, "pass_word": {passWord}})
	defer rsp.Body.Close()
	if err != nil {
		return rsp.StatusCode, nil, err
	}
	data, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return rsp.StatusCode, nil, err
	}
	return rsp.StatusCode, data, err
}

//发起注册请求
func ReqRegister(userName string, passWord string) (code int, data []byte, err error) {
	rsp, err := http.PostForm("http://127.0.0.1:8080/v1/user/register",
		url.Values{"user_name": {userName}, "pass_word": {passWord}})
	defer rsp.Body.Close()
	if err != nil {
		return rsp.StatusCode, nil, err
	}
	data, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return rsp.StatusCode, nil, err
	}
	return rsp.StatusCode, data, err
}
