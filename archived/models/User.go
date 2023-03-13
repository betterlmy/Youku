package models

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	"reflect"
	"time"
)

type URL = string

type User struct {
	Id      int
	Name    string
	Status  int
	Mobile  string
	AddTime int64
	Avatar  string
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
	orm.RegisterModel(new(User)) // 注册模型,因为orm需要知道你要操作的表 ORM可以自动根据struct创建表
}

// GetUserInfo 获取用户信息
func GetUserInfo(id int) (u User, err error) {
	o := orm.NewOrm()
	u = User{Id: id}
	err = o.Read(&u)
	return

}

// RegisterNewUser 注册新用户 返回用户信息
func RegisterNewUser(name, mobile string) (user1 User, err error) {
	o := orm.NewOrm()
	user1 = User{
		Name:    name,
		Avatar:  "default avatar path",
		Mobile:  mobile,
		Status:  0,
		AddTime: time.Now().Unix(), // int64类型的时间戳timestamp
		//AddTime: time.Now().Format("2006-01-02 15:04:05"), // 其他日期格式 string
	}

	tableLength, err := o.Insert(&user1)
	// sqlRaw
	//sql := "INSERT INTO user (id, name, status, mobile, avatar, add_time) VALUES (?,?,?,?,?,?)"
	//_, err = o.Raw(sql, 0, name, 0, mobile, user1.Avatar, user1.AddTime).Exec()
	_ = tableLength
	return
}

// UpdateUserInfo 根据id name avatar选择性更新用户信息
func UpdateUserInfo(id int, name, avatar string) (u User, err error) {
	o := orm.NewOrm()
	u = User{Id: id}
	err = o.Read(&u)
	if err != nil {
		return
	}
	if name != "" {
		// 可以添加别的限制条件
		// u.Name =check(name)
		u.Name = name
	}
	if avatar != "" {
		u.Avatar = avatar
	}
	sql := "UPDATE user SET name=?,avatar=? WHERE id=?"
	_, err = o.Raw(sql, u.Name, u.Avatar, u.Id).Exec()
	tableLength, err := o.Update(&u)
	_ = tableLength
	return

}

// DeleteUser 根据id删除用户
func DeleteUser(id int) (err error) {
	o := orm.NewOrm()

	sql := "DELETE FROM user WHERE id=?"
	_, err = o.Raw(sql, id).Exec()
	_, err = o.Delete(&User{Id: id})
	// 如果id是[]int 类型 传入的参数是[]int类型 例如 []int{1,2,3} 会删除id为1,2,3的用户
	// 如果id是int类型 传入的参数是int类型 例如 1 会删除id为1的用户
	// 实现一次性删除多个用户的方法
	return
}

// List 获取用户列表
func List() (users []User, err error) {
	o := orm.NewOrm()          // 获取orm对象 orm对象是操作数据库的对象 一个orm对象对应一个数据库 一个数据库可以有多个orm对象 一个orm对象可以操作多个表 一个表只能有一个orm对象
	qs := o.QueryTable("user") // 获取查询集对象 一个查询集对象对应一张表
	qs = qs.Filter("id__gt", 1)
	// 过滤条件 gt大于  lt小于  gte大于等于  lte小于等于  in包含  contains包含  startswith以什么开头  endswith以什么结尾  exact精确匹配  iexact不区分大小写精确匹配
	qs = qs.Limit(10)      // 限制查询结果数量 2条 0表示不限制 默认0
	qs = qs.OrderBy("-id") // 排序 -id倒序  id正序 默认正序
	// 可以多个字段排序 逗号分隔 例如 qs.OrderBy("-id","name") 优先按id倒序排序 再按name正序排序
	// 打印sql语句

	_, err = qs.All(&users, "Id", "Name") // 查询所有数据 传入一个切片指针 传入需要查询的字段名
	//只有执行.All()、.One()、.Values()等方法时，ORM系统才会根据这些参数和条件生成相应的SQL查询并执行，最终返回查询结果。

	sql := "SELECT * FROM user WHERE id > ? ORDER BY id DESC LIMIT 10"
	_, err = o.Raw(sql, 1).QueryRows(&users)
	if err != nil {
		return nil, err
	} // 使用Raw方法来拼接id,而非直接将id写入sql语句中，防止sql注入

	if err != nil {
		return nil, err
	}
	return

}

// SqlUserInfo 通过sql语句来获取用户信息
func SqlUserInfo(id int) (u User, err error) {
	o := orm.NewOrm()
	sql := "SELECT * FROM user WHERE id = ? LIMIT 1 "
	err = o.Raw(sql, id).QueryRow(&u) // 使用Raw方法来拼接id,而非直接将id写入sql语句中，防止sql注入
	// QueryRow方法只能查询一条数据，如果查询多条数据会报错
	return
}
