package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Uid      int     `gorm:"size:64;not null;unique;primary_key"` //uid值，长度 255, 不为空，唯一值, 主键
	Uname    string  `gorm:"type:varchar(40);not null;unique"`
	Password string  `gorm:"type:varchar(128)"`
	Status   int     `gorm:"size:4"`
	Group    []Group `gorm:"many2many:user_groups"`
}

type Group struct {
	Gid    int    `gorm:"size:64;unique;primary_key"`
	Name   string `gorm:"type:varchar(64);unique"`
	Status int    `gorm:"size:4"`
}

type User_group struct {
	Status int `gorm:"size:4"`
}

type Group_relation struct {
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/tests2")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&User{}, &Group{}, &User_group{})
}
