package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["fyoukuApi/controllers:UserController"] = append(beego.GlobalControllerRouter["fyoukuApi/controllers:UserController"],
		beego.ControllerComments{
			Method:           "SaveRegister",
			Router:           "/register/save",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}