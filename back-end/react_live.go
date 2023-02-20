package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func getOrigin() *url.URL {
	origin, _ := url.Parse("http://localhost:3000")
	return origin
}

var origin = getOrigin()

var director = func(req *http.Request) {
	req.Header.Add("X-Forwarded-Host", req.Host)
	req.Header.Add("X-Origin-Host", origin.Host)
	req.URL.Scheme = "http"
	req.URL.Host = origin.Host
}

// ReactHandler loads angular assets
var ReactHandler = &httputil.ReverseProxy{Director: director}
