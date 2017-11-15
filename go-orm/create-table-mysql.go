package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type user struct {
	uid      int    `gorm:"size:64;not null;unique;primary_key"` //uid值，长度 255, 不为空，唯一值, 主键
	uname    string `gorm:"size:l40"`
	password string `gorm:"size:128`
	status   int    `gorm:"size:4"`
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/datrix")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&user{})

}
