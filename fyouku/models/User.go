package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
)

type UserInfo struct {
	Id      int
	Name    string
	AddTime int64
	Avatar  string
}

// 判断用户名密码是否正确
func IsMobileLogin(mobile string, password string) string {
	req := httplib.Post(getApiURL() + "/login/do")
	//req := httplib.Post(beego.AppConfig.String("microApi") + "/fyoukuApi/user/user/LoginDo")
	req.Param("mobile", mobile)
	req.Param("password", password)
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	return str
}

// 用户注册
func UserSave(mobile, password string) string {
	req := httplib.Post(getApiURL() + "/register/save")
	req.Param("mobile", mobile)
	req.Param("password", password)
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	return str
}

// 保存用户
func SendMessageDo(uids string, content string) string {
	req := httplib.Post(getApiURL() + "/send/message")
	req.Param("uids", uids)
	req.Param("content", content)
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	return str
}
