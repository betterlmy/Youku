package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "youku/routers"
)

func main() {
	//beego.SetStaticPath("/download", "download")
	defaultdb := beego.AppConfig.String("defaultdb")
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		fmt.Println(err)
		panic("驱动注册失败")
	}
	// orm必须注册一个别名为default的数据库 否则编译无法通过
	err = orm.RegisterDataBase("default", "mysql", defaultdb)
	if err != nil {
		panic("数据库注册失败")
	}

	beego.Run()
}
