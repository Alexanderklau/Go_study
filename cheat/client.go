// client
package main

import (
	"fmt"
	"net"
	"os"
)

var ch chan int = make(chan int)
var nickname string

func reader(conn *net.TCPConn) {
	buff := make([]byte, 256)
	for {
		j, err := conn.Read(buff)
		if err != nil {
			ch <- 1
			break
		}
		fmt.Printf("%s\n", buff[0:j])
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(os.Stderr, "Usage:%s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	TcpAdd, _ := net.ResolveTCPAddr("tcp", service)
	conn, err := net.DialTCP("tcp", nil, TcpAdd)
	if err != nil {
		fmt.Println("服务没打开")
		os.Exit(1)
	}
	defer conn.Close()
	go reader(conn)
	fmt.Println("请输入昵称")
	fmt.Scanln("你的昵称是:", nickname)
	for {
		var msg string
		fmt.Scan(&msg)
		fmt.Print("<" + nickname + ">" + "说：")
		// }
		fmt.Println(msg)
		b := []byte("<" + nickname + ">" + "说：" + msg)
		conn.Write(b)
		select {
		case <-ch:
			fmt.Println("server错误")
			os.Exit(2)
		default:
		}
	}
}
