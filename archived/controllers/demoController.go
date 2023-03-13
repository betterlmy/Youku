package controllers

import (
	"github.com/beego/beego/v2/adapter/logs"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
	"strings"
	"time"
	"youku/models"
)

type DemoController struct {
	beego.Controller
}

/*
ORM CRUD
*/

// 通过id获取用户信息  Retrieve功能
// @router /user/getInfo [get]
func (c *DemoController) GetUsername() {
	var (
		id    int
		err   error
		title string
		user  models.User
	)
	// 接收浏览器的参数
	id, err = c.GetInt("id")

	if err != nil {
		logs.Error("id不存在")
		c.Ctx.WriteString("id不存在,获取失败")
		return
	}
	logs.Info("查询id:%d", id)

	// 从数据库获取信息
	user, err = models.GetUserInfo(id)
	out := ""
	if err != nil {
		title = "数据库查询失败"
		out = title
	} else {
		title = "获取成功"
		out = title + "\n" + "用户名:" + user.Name
		out += "\n添加时间:" + time.Unix(user.AddTime, 0).Format("2006-01-02 15:04:05")
	}
	// 前端展示
	c.Ctx.WriteString(out)
}

// 用户注册功能  Create功能
// @router /user/register [get]
func (c *DemoController) Register() {
	// 需求 姓名 手机号 id 自增 其他自动获取或者置空
	var (
		name   string
		mobile string
		user   models.User
		err    error
	)

	// 使用get请求获取字段
	name = c.GetString("name")
	mobile = c.GetString("mobile")
	if name == "" {
		logs.Error("name字段为空,添加失败")
		c.Ctx.WriteString("name字段为空,添加失败")
		return
	}
	if mobile == "" {
		logs.Error("mobile字段为空,添加失败")
		c.Ctx.WriteString("mobile字段为空,添加失败")
		return
	}
	if err != nil {
		logs.Error(err)
		c.Ctx.WriteString("添加失败")
	} else {
		// 调用models的注册方法
		user, err = models.RegisterNewUser(name, mobile)
		// 输出新添加的用户信息
		c.Ctx.WriteString("添加成功,新用户信息如下:\n")
		c.Ctx.WriteString(user.Print())

	}

}

// 修改功能   Update功能
// @router /user/update [get]
func (c *DemoController) Update() {
	// 根据get的id,修改其它参数
	var (
		id     int
		name   string = ""
		avatar string = ""
		err    error
	)

	// 获取参数
	id, err = c.GetInt("id")

	if err != nil {
		logs.Error("id不存在")
		c.Ctx.WriteString("id不存在,获取失败")
		return
	}
	logs.Info("查询id:%d", id)

	name = c.GetString("name")
	avatar = c.GetString("avatar")
	user, err := models.UpdateUserInfo(id, name, avatar)
	if err != nil {
		logs.Error(err)
		c.Ctx.WriteString("修改失败")
		return
	} else {
		c.Ctx.WriteString("修改成功,修改后的用户信息如下:\n")
		c.Ctx.WriteString(user.Print())
	}
}

// 删除功能(可同时指定多个id)   Delete功能
// @router /user/delete [get]
func (c *DemoController) Delete() {
	// 根据get的id,删除用户
	var (
		id      int
		err     error
		outputs []string
	)
	// 获取参数
	idsStr := c.GetString("id")
	ids := strings.Split(idsStr, ",")

	// 遍历删除
	for _, v := range ids {
		id, err = strconv.Atoi(v)
		if err != nil {
			logs.Error("id字段转换失败")
			continue
		}
		err = models.DeleteUser(id)
		if err != nil {
			logs.Error(err)
			logs.Info("删除id:%d失败", id)
			outputs = append(outputs, string(id)+"删除失败")
			continue
		} else {
			logs.Info("删除id:%d成功", id)
			outputs = append(outputs, string(id)+"删除成功")
		}
	}

	// 输出前端
	output := ""
	for _, v := range outputs {
		output += v + "\n"
	}
	c.Ctx.WriteString(output)
}

// 获取用户列列表 高级查询
// @router /user/list [get]
func (c *DemoController) GetUsers() {
	// 从数据库中返回所有的用户数量等等
	var (
		users  []models.User
		err    error
		output string
	)
	users, err = models.List()
	if err != nil {
		logs.Error(err)
		return
	}
	if len(users) == 0 {
		c.Ctx.WriteString("用户列表为空")
		return
	}
	for _, user := range users {
		output += user.Print()
		output += "------------\n"
	}
	c.Ctx.WriteString(output)
}

// 通过id获取用户信息 SqlRaw代码
// @router /user/sql/getInfo [get]
func (c *DemoController) SqlGetUserInfo() {
	var (
		id    int
		err   error
		title string
		user  models.User
	)
	// 接收浏览器的参数
	id, err = c.GetInt("id")

	if err != nil {
		logs.Error("id不存在")
		c.Ctx.WriteString("id不存在,获取失败")
		return
	}
	logs.Info("查询id:%d", id)

	// 从数据库获取信息
	user, err = models.SqlUserInfo(id)
	out := ""
	if err != nil {
		title = "数据库查询失败"
		out = title
	} else {
		title = "获取成功"
		out = title + "\n" + "用户名:" + user.Name
		out += "\n添加时间:" + time.Unix(user.AddTime, 0).Format("2006-01-02 15:04:05")
	}
	// 前端展示
	c.Ctx.WriteString(out)
}
