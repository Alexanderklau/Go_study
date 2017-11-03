package main

import "fmt"

func main() {
	const (
		a = iota //0
		b        //1
		c        //2
		k
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, k, d, e, f, g, h, i)

}
