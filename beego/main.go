package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("root:root@tcp(127.0.0.1:3306)/datrix2?charset=utf8")
}

func main() {
	o := orm.NewOrm()
	o.Using("datrix2") // 默认使用 default，你可以指定为其他数据库

	profile := new(Profile)
	profile.Age = 30

	user := new(User)
	user.Profile = profile
	user.Name = "slene"

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))
}
