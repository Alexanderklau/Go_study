package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) string {
	fmt.Println(r.Host)
	// output: localhost:9090
	fmt.Println(r.RequestURI)
	// output: /index?id=1
	c := r.RequestURI
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	fmt.Println(strings.Join([]string{scheme, r.Host, r.RequestURI}, ""))
	// output: http://localhost:9090/index?id=1
	return c
}

func main() {
	http.HandleFunc("/", index)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
