package controllers

import beego "github.com/beego/beego/v2/server/web"

type DemoController struct {
	beego.Controller
}

// @router /hello1 [get]
func (c *DemoController) GetHello() {
	c.Ctx.WriteString("Hello World")
}

// ORM CRUD

// 通过id获取用户名
