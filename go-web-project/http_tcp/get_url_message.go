package main

import (
	"fmt"
	"log"
	"net/http"
)

//http://10.0.9.127:9090/?fileidstr=a0a1adc1f7b4d8ad58e2a396e70bcf8d.xlsx&iswindows=1&optuser=admin

func addUser(w http.ResponseWriter, req *http.Request) {
	fileId := req.FormValue("fileidstr")
	user := req.FormValue("optuser")
	out := "http://10.0.9.139/viewer/dcomp.php?fileidstr=" + fileId + "&optuser=" + user
	fmt.Fprintf(w, out)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!")
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/user", addUser)
	err := http.ListenAndServe(":8880", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
