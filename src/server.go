package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type server struct {
	URL          string
	ReverseProxy *httputil.ReverseProxy
	Health       bool
}

func newServer(urlStr string) *server {
	u, _ := url.Parse(urlStr)
	rp := httputil.NewSingleHostReverseProxy(u)
	return &server{
		URL:          urlStr,
		ReverseProxy: rp,
		Health:       true,
	}
}

func (s *server) checkHealth() {

	res, err := http.Head(s.URL)

	if err != nil {
		log.Println(err)
	}

	if res.StatusCode != http.StatusOK {
		s.Health = false
	} else {
		s.Health = true
	}

	return
}
