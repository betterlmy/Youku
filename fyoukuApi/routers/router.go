package routers

import (
	"fyoukuApi/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Include(&controllers.UserController{})
}
