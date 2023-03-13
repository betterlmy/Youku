package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"regexp"
)

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Item  interface{} `json:"item"`
	Count int64       `json:"count"`
}

// ReturnSuccess 正确的返回值 传入状态码和信息 返回JsonStruct 用于前端展示
func ReturnSuccess(code int, msg, item interface{}, count int64) *JsonStruct {
	return &JsonStruct{
		Code:  code,
		Msg:   msg,
		Item:  item,
		Count: count,
	}
}

// ReturnError 错误的返回值 传入错误码和错误信息 返回JsonStruct 用于前端展示
func ReturnError(code int, msg interface{}) *JsonStruct {
	return &JsonStruct{
		Code: code,
		Msg:  msg,
	}
}

// checkPhoneFmt 判断手机号格式是否正确
func checkPhoneFmt(mobile string) bool {
	matched, _ := regexp.MatchString(`^1[3-9]\d{9}$`, mobile)
	return matched
}

// MD5V 将密码通过MD5拼接一个字符串进行加密操作
func MD5V(password string) string {
	suffix := beego.AppConfig.String("md5code")
	if suffix == "" {
		logs.Error("缺少md5code配置")
		suffix = "youku"
	}
	password += suffix
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}
