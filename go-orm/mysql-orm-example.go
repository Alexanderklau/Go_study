package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:123456@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("数据库连接失败")
	}
	defer db.Close()
}
