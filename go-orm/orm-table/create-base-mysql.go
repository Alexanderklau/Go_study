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
	"./structure" //引入包structure
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zheng-ji/goSnowFlake"
	"strconv"
	_ "time"
)

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/datrix2")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&structure.User{}, &structure.Group{}, &structure.User_group{})
	iw, err := goSnowFlake.NewIdWorker(1)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 10; i++ {
		if id, err := iw.NextId(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(id)
			user := structure.User{Uid: int(id), Uname: "lwb" + strconv.Itoa(i), Password: "123456", Status: 1}
			group := structure.Group{Gid: int(id), Name: "datatom" + strconv.Itoa(i), Status: 1}
			fmt.Println(group)
			db.NewRecord(group)
			db.Create(&group)
			db.NewRecord(user)
			db.Create(&user)
		}
	}
}
