package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "betterlmy.com"
	c.Data["Email"] = "betterlmy@icloud.com"
	c.Data["IsEmail"] = 0
	pages := []struct {
		Num int
	}{{10}, {20}, {30}}
	c.Data["Pages"] = pages
	c.TplName = "index.tpl"

}

func (c *MainController) GetHello() {
	c.Ctx.WriteString("Hello Beego!")
}
