package controllers

import "github.com/astaxie/beego"

type DemoController struct {
	beego.Controller
}

// 输出HelloWorld
// @router /demo/hello [get]
func (c *DemoController) GetHello() {
	c.Ctx.WriteString("Hello World")

}
