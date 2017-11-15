package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:123456@/datrix")
	if err != nil {
		panic("数据库连接失败")
	}
	defer db.Close()
}
