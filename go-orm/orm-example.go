package main

import (
	"github.com/jinzhu/gorm"
	_ "github.github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("数据库链接失败！")
	}
	defer db.Close()

	db.AutoMigreate(&Product{})

	db.Create(&Product{Code: "L1212", Price: 1000})

	var product Product

	db.First(&product, 1)
}
