package main

// http://10.0.7.96/hosting/discovery

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://10.0.7.96/hosting/discovery")
	if err != nil {
		fmt.Println("Http get error!")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error.")
	}
	src := string(body)
	re, _ := regexp.Compile("关于.{1,2}")
	src = re.FindString(src)
	fmt.Println(src)
}
