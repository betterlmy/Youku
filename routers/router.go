package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"youku/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &controllers.MainController{}, "get:GetHello")

	beego.InsertFilter(
		"/demo/*",
		beego.BeforeRouter,
		FilterDemo,
	)
}

func FilterDemo(ctx *context.Context) {
	ctx.WriteString("禁止访问")
}
