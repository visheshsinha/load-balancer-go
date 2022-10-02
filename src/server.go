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
	serverList = []*httputil.ReverseProxy{
		createHost("http://127.0.0.1:5000"),
		createHost("http://127.0.0.1:5001"),
		createHost("http://127.0.0.1:5002"),
		createHost("http://127.0.0.1:5003"),
		createHost("http://127.0.0.1:5004"),
	}
	lastServerIndex = 0
)

func forwardRequest(res http.ResponseWriter, req *http.Request) {
	// fmt.Fprintln(res, "Hello from Load-Balancer")
	server := getServer()
	server.ServeHTTP(res, req)
}

func getServer() *httputil.ReverseProxy {
	nextIndex := (lastServerIndex + 1) % len(serverList)
	server := serverList[nextIndex]
	lastServerIndex = nextIndex
	return server
}

func createHost(urlStr string) *httputil.ReverseProxy {
	url, _ := url.Parse(urlStr)
	log.Printf("Creating Reverse Proxy for the URL: %s", url.String())
	return httputil.NewSingleHostReverseProxy(url)
}