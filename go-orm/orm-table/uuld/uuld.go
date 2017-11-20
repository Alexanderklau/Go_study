package main

import (
	"../structure"
	"fmt"
	"github.com/zheng-ji/goSnowFlake"
	"strconv"
)

func main() {
	// Params: Given the workerId, 0 < workerId < 1024
	iw, err := goSnowFlake.NewIdWorker(1)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 10; i++ {
		if id, err := iw.NextId(); err != nil {
			fmt.Println(err)
		} else {
			st := structure.User{Uid: int(id), Uname: "lwb" + strconv.Itoa(i), Password: "123456", Status: 1}
		}
	}
}
