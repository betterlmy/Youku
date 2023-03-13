package models

import (
	beego "github.com/beego/beego/v2/server/web"
)

func getApiURL() string {
	url, _ := beego.AppConfig.String("apiurl")
	return url
}
