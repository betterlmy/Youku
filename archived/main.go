package main

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	_ "youku/routers"
)

func dbRegister(defaultdb string) (err error) {
	err = orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		return err
	}
	// orm必须注册一个别名为default的数据库 否则编译无法通过
	err = orm.RegisterDataBase("default", "mysql", defaultdb)
	orm.Debug = true
	return err
}

func main() {
	//beego.SetStaticPath("/download", "download")
	defaultdb, _ := beego.AppConfig.String("defaultdb")
	err := dbRegister(defaultdb)

	if err != nil {
		panic(err)
	}
	logs.Info(defaultdb, "数据库已连接")
	beego.Run()
}
