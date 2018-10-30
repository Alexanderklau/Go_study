package main

import "fmt"
import "io/ioutil"

func main() {
	dir_list, e := ioutil.ReadDir("/home/lau/下载")
	if e != nil {
		fmt.Println("read dir error")
		return
	}
	for i, v := range dir_list {
		fmt.Println(i, "=", v.Name())
	}
}
