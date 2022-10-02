package main

import (
	"log"
	"net/url"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", forwardRequest)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

var (
	serverList = []string{
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5001",
		"http://127.0.0.1:5002",
		"http://127.0.0.1:5003",
		"http://127.0.0.1:5004",
	}
	nextIndex = 0
	lastServerIndex = 0
)

func forwardRequest(res http.ResponseWriter, req *http.Request) {
	// fmt.Fprintln(res, "Hello from Load-Balancer")
	url := getServer()
	rproxy := httputil.NewSingleHostReverseProxy(url)
	log.Printf("Routing the request to the URL: %s", url.String())
	rproxy.ServeHTTP(res, req)
}

func getServer() *url.URL {
	nextIndex = (lastServerIndex + 1) % 5
	url, _ := url.Parse(serverList[nextIndex])
	lastServerIndex = nextIndex
	return url
}