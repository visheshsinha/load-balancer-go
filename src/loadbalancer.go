package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", forwardRequest)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

var (
	serverList = []*server{
		newServer("http://127.0.0.1:5001"),
		newServer("http://127.0.0.1:5002"),
		newServer("http://127.0.0.1:5003"),
		newServer("http://127.0.0.1:5004"),
		newServer("http://127.0.0.1:5005"),
	}
	lastServerIndex = 0
	nextIndex = 0
)

func forwardRequest(res http.ResponseWriter, req *http.Request) {
	server := getServer()
	server.ReverseProxy.ServeHTTP(res, req)
}

func getServer() *server {
	server := serverList[nextIndex]
	nextIndex = (lastServerIndex + 1) % len(serverList)
	lastServerIndex = nextIndex
	return server
}