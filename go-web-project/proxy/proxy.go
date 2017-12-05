package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	// "strings"
)

type handle struct {
	host string
	port string
}

func newReverseProxy(target *url.URL) *httputil.ReverseProxy {
	targetQuery := target.RawQuery
	director := func(req *http.Request) {
		req.Host = target.Host
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		// req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		// if targetQuery == "" || req.URL.RawQuery == "" {
		// 	req.URL.RawQuery = targetQuery + req.URL.RawQuery
		// } else {
		// 	req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		// }
		// if _, ok := req.Header["User-Agent"]; !ok {
		// 	// explicitly disable User-Agent so it's not set to defaul t value
		// 	req.Header.Set("User-Agent", "")
		// }
	}
	return &httputil.ReverseProxy{Director: director}
}

// func singleJoiningSlash(a, b string) string {
// 	aslash := strings.HasSuffix(a, "/")
// 	bslash := strings.HasPrefix(b, "/")
// 	switch {
// 	case aslash && bslash:
// 		return a + b[1:]
// 	case !aslash && !bslash:
// 		return a + "/" + b
// 	}
// 	return a + b
// }

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse("http://" + this.host + ":" + this.port)
	log.Println(remote)
	if err != nil {
		panic(err)
	}
	proxy := newReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

func startServer() {
	//被代理的服务器host和port
	h := &handle{host: "10.0.7.96", port: "80"}
	log.Println(h)
	err := http.ListenAndServe(":8090", h)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

func main() {
	startServer()
}
