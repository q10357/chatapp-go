package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	microserviceURL, _ := url.Parse("http://localhost:8080/") // URL of the microservice

	// Setup the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(microserviceURL)

	// Setup the handler to our server to accept incoming connections
	http.HandleFunc("/", handleRequestAndRedirect)

	// Start the server
	log.Fatal(http.ListenAndServe("127.0.0.1:8081", proxy)) // Gateway listens on port 8081
}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s %s\n", req.Method, req.URL, req.RemoteAddr)

	// Check if the request is for the 'authservice'
	if strings.HasPrefix(req.URL.Path, "/auth") {
		req.URL.Host = "localhost:8080"
		req.URL.Scheme = "http"
		req.URL.Path = strings.TrimPrefix(req.URL.Path, "/auth")
	} else {
		// Default service
		req.URL.Host = "localhost:8081"
		req.URL.Scheme = "http"
	}

	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = req.URL.Host
}
