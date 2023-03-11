package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"youku/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &controllers.MainController{}, "get:GetHello")
	beego.Include(&controllers.DemoController{})
	beego.InsertFilter(
		"/demo/*",
		beego.BeforeRouter,
		FilterDemo,
	)
}

func FilterDemo(ctx *context.Context) {
	ctx.WriteString("禁止访问")
}
