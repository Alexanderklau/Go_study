package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest(r.Method, r.RequestURI, r.Body)
	for k, v := range r.Header {
		for _, vv := range v {
			req.Header.Add(k, vv)
		}
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("error:", err)
	}

	defer resp.Body.Close()

	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	for _, value := range resp.Request.Cookies() {
		w.Header().Add(value.Name, value.Value)
	}

	w.WriteHeader(resp.StatusCode)

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
	}

	_, err = w.Write(result)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Start serving on port 9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Println("error:", err)
	}
}
