package main

import (
	_ "fyoukuApi/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //很重要
)

func dbRegister(defaultdb string) (err error) {
	err = orm.RegisterDriver("mysql", orm.DRMySQL)
	logs.Info("数据库驱动注册成功")
	if err != nil {
		return err
	}
	// orm必须注册一个别名为default的数据库 否则编译无法通过
	err = orm.RegisterDataBase("default", "mysql", defaultdb)
	orm.Debug = true
	return err
}

func main() {
	defaultdb := beego.AppConfig.String("defaultdb")
	err := dbRegister(defaultdb)

	if err != nil {
		panic(err)
	}
	logs.Info(defaultdb, "数据库已连接")

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
