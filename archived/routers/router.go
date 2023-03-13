package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"youku/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/hello", &controllers.MainController{}, "get:GetHello")

	beego.Include(&controllers.DemoController{})
	beego.InsertFilter(
		// 过滤器的作用范围
		"/demo/*",
		beego.BeforeRouter, // 在路由之前执行
		FilterDemo,         // 满足过滤条件执行的函数
	)
}

func FilterDemo(ctx *context.Context) {
	ctx.WriteString("禁止访问")
}
