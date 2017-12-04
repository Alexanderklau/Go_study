package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	fmt.Println(req.RequestURI)
	fmt.Println(req.URL)
	w.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":80", nil)
}
