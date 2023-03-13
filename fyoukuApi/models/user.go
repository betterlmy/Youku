package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"reflect"
	"time"
)

type User struct {
	Id       int    `orm:"pk;auto"`
	Name     string `orm:"size(20),default('default_name')"`
	Mobile   string `orm:"size(11)"`
	Avatar   string `orm:"size(255),default('icon/default_avatar.png')"`
	Status   int    `orm:"default(0)"`
	AddTime  int64  `orm:"int(11)"`
	Password string `orm:"size(255)"`
}

func (u User) Print() (output string) {
	v := reflect.ValueOf(u)
	// 遍历结构体字段
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		output += fmt.Sprintf(" %v:%v\n", v.Type().Field(i).Name, fieldValue.Interface())
	}
	return
}

func init() {
	orm.RegisterModel(new(User))
}

// IsUserMobile 判断用户是否存在 存在则返回true 不存在则返回false
func IsUserMobile(mobile string) bool {
	o := orm.NewOrm()
	user := User{Mobile: mobile}
	err := o.Read(&user, "mobile") // 通过mobile字段查询
	if err == nil {
		// 存在
		return true
	}
	return false
}

// UserSave 保存用户信息
func UserSave(mobile, password string) (err error) {
	o := orm.NewOrm()
	user := User{Mobile: mobile,
		Password: password,
		AddTime:  time.Now().Unix(),
	}
	logs.Info(user.Print())
	x, err := o.Insert(&user)
	logs.Info(x, err)
	return
}
