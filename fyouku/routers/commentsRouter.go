package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["fyouku/controllers:MainController"] = append(beego.GlobalControllerRouter["fyouku/controllers:MainController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/index`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["fyouku/controllers:TestController"] = append(beego.GlobalControllerRouter["fyouku/controllers:TestController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/test/index`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
