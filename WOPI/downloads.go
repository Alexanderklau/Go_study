package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	url := "http://211.144.114.26:15300/viewer/dcomp.php?fileidstr=378b25d25189e113887ed9862ed7c10e.wmv&optuser=test"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("11.wmv")
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
}
