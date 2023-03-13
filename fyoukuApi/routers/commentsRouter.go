package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["fyoukuApi/controllers:UserController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:UserController"],
		beego.ControllerComments{
			Method:           "SaveRegister",
			Router:           `/register/save`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
