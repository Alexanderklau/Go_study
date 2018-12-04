package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type configuration struct {
	Ip       string
	Log_type string
	Port     string
	Enabled  bool
	Path     string
}

func main() {
	//打开文件
	file, _ := os.Open("config.json")

	//关闭文件
	defer file.Close()

	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)

	conf := configuration{}
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 写入日志
	log_file := conf.Path

	// 读取文件
	files, _ := ioutil.ReadFile(log_file)
	if err != nil {
		fmt.Println(err)
	}

	// 输出文件
	str := string(files)
	fmt.Println(str)
}
