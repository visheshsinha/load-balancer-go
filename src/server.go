package main

import (
	"net/http/httputil"
	"net/url"
)

type server struct {
	URL          string
	ReverseProxy *httputil.ReverseProxy
	Health       bool
}

func newServer(urlStr string) *server{
	u, _ := url.Parse(urlStr)
	rp := httputil.NewSingleHostReverseProxy(u)
	return &server{
		URL:          urlStr,
		ReverseProxy: rp,
		Health:       true,
	}
}
