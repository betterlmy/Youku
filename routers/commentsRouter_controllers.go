package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["youku/controllers:DemoController"] = append(beego.GlobalControllerRouter["youku/controllers:DemoController"],
		beego.ControllerComments{
			Method:           "GetHello",
			Router:           "/hello1",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
