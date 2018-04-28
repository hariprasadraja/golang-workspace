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

func RedirectToSubDomine(w http.ResponseWriter, r *http.Request) {
	defer panicRecovery()

	parts := strings.Split(r.Host, ".")
	if len(parts) != 3 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// canonical name of the Domine.
	redirectUrl := "http:localhost:8080"
	targetUrl := &url.URL{Scheme: "http", Host: redirectUrl}
	proxy := httputil.NewSingleHostReverseProxy(targetUrl)
	proxy.ServeHTTP(customWriter{w}, r)
}

type customWriter struct {
	http.ResponseWriter
}

func (w customWriter) WriteHeader(c int) {
	if c == http.StatusBadGateway {
		http.Error(w.ResponseWriter, "Server under maintenance", c)
		return
	}

	w.ResponseWriter.WriteHeader(c)
}

func (w customWriter) Write(b []byte) (int, error) {
	return w.ResponseWriter.Write(b)
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

	go func() {
		sslmux := http.NewServeMux()
		sslmux.HandleFunc("/", RedirectToSubDomine)

		crtFile := ""
		keyFile := ""

		tlsconf := &tls.Config{ServerName: ""}
		sslServer := http.Server{
			Addr:         ":443",
			ReadTimeout:  90 * time.Second,
			WriteTimeout: 90 * time.Second,
			Handler:      sslmux,
			TLSConfig:    tlsconf,
		}

		log.Print("Listening on: 443")
		err := sslServer.ListenAndServeTLS(crtFile, keyFile)
		if err != nil {
			log.Printf("Error in running ssl proxy server: %s", err.Error())
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/", redirectToSSLServer)

	server := http.Server{
		Addr:         ":90",
		ReadTimeout:  90 * time.Second,
		WriteTimeout: 90 * time.Second,
		Handler:      mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Error in proxy server: %+v", err.Error())
	}

	defer server.Close()
	log.Printf("listerning on port : %s", server.Addr)

}
