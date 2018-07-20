package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime/debug"
	"strings"
	"time"
)

func panicRecovery() {
	if err := recover(); err != nil {
		log.Printf("panic: %+v", err)
		debug.PrintStack()
	}
}

// RedirectToSubDomine redirects the request to sudomine server
func RedirectToSubDomine(w http.ResponseWriter, r *http.Request) {
	defer panicRecovery()

	parts := strings.Split(r.Host, ".")
	if len(parts) != 3 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	targetURL := &url.URL{Scheme: "http", Host: "http:localhost:8080"} //cannonical name
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.ServeHTTP(w, r)
}

func redirectToSSLServer(w http.ResponseWriter, r *http.Request) {
	defer panicRecovery()

	parts := strings.Split(r.Host, ".")
	if len(parts) != 3 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	urlstr := fmt.Sprintf("https://%s%s", r.Host, r.URL.String())
	http.Redirect(w, r, urlstr, http.StatusMovedPermanently)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", redirectToSSLServer)
	server := http.Server{
		Addr:         ":9000",
		ReadTimeout:  200 * time.Second,
		WriteTimeout: 200 * time.Second,
		Handler:      mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Error in proxy server: %+v", err.Error())
	}

	defer server.Close()
	log.Printf("listerning on port : %s", server.Addr)
	go runSSlServer()
}

func runSSlServer() {
	sslmux := http.NewServeMux()
	sslmux.HandleFunc("/", RedirectToSubDomine)
	crtFile := ""
	keyFile := ""
	tlsconf := &tls.Config{ServerName: ""}
	sslServer := http.Server{
		Addr:         ":443",
		ReadTimeout:  200 * time.Second,
		WriteTimeout: 200 * time.Second,
		Handler:      sslmux,
		TLSConfig:    tlsconf,
	}

	log.Print("Listening on: 443")
	err := sslServer.ListenAndServeTLS(crtFile, keyFile)
	if err != nil {
		log.Printf("Error in running ssl proxy server: %s", err.Error())
	}
}
