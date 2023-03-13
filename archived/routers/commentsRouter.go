package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["youku/controllers:DemoController"] = append(beego.GlobalControllerRouter["youku/controllers:DemoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/user/delete`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["youku/controllers:DemoController"] = append(beego.GlobalControllerRouter["youku/controllers:DemoController"],
		beego.ControllerComments{
			Method:           "GetUsername",
			Router:           `/user/getInfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["youku/controllers:DemoController"] = append(beego.GlobalControllerRouter["youku/controllers:DemoController"],
		beego.ControllerComments{
			Method:           "GetUsers",
			Router:           `/user/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["youku/controllers:DemoController"] = append(beego.GlobalControllerRouter["youku/controllers:DemoController"],
		beego.ControllerComments{
			Method:           "Register",
			Router:           `/user/register`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["youku/controllers:DemoController"] = append(beego.GlobalControllerRouter["youku/controllers:DemoController"],
		beego.ControllerComments{
			Method:           "SqlGetUserInfo",
			Router:           `/user/sql/getInfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["youku/controllers:DemoController"] = append(beego.GlobalControllerRouter["youku/controllers:DemoController"],
		beego.ControllerComments{
			Method:           "Update",
			Router:           `/user/update`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
