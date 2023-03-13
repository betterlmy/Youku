package controllers

import (
	"fyoukuApi/models"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

// SaveRegister 用户注册功能
// @router /register/save [post]
func (c *UserController) SaveRegister() {
	logs.Info("尝试注册")
	var (
		mobile   string
		password string
	)

	// 接收参数
	mobile = c.GetString("mobile")
	password = c.GetString("password")
	logs.Info("mobile:", mobile, "password:", password)
	// 判断格式
	if mobile == "" {
		c.Data["json"] = ReturnError(4001, "手机号不能为空") // 将json数据返回给前端
	} else if password == "" {
		c.Data["json"] = ReturnError(4003, "密码不能为空")
	} else if !checkPhoneFmt(mobile) {
		c.Data["json"] = ReturnError(4002, "手机号格式错误")
	} else if models.IsUserMobile(mobile) {
		c.Data["json"] = ReturnError(4004, "手机号已注册")
	} else {
		// 保存到数据库
		err := models.UserSave(mobile, MD5V(password))
		if err != nil {
			c.Data["json"] = ReturnError(4005, "注册失败")
		} else {
			c.Data["json"] = ReturnSuccess(0, "注册成功", nil, 0)
		}

	}
	logs.Info(c.Data["json"])
	_ = c.ServeJSON() // 框架将json个数的数据返回给前端
}
