package main

/*
生成基本的表，user,group,group_relation,等
如果要自己添加字段，在相应的模型下面定义字段名，类型，特征等
example：在user下添加create_time字段，类型为时间
create_time   time.Time
运行程序以后会自动构建相应的表，请注意，如果添加了不需要的字段，删除之后再次构建字段还是会存在的
这时候需要运行删除字段的命令
*/
import (
	"./structure"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "time"
)

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/datrix2?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&structure.User{}, &structure.Group{})

}
