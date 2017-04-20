package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/justinas/alice"
)

type numberDumper int

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ServeHttp")
	fmt.Fprintf(w, "Here's your number: %d\n", n)
	a:="hello buddy"
	b:= []byte(a)
	w.Write(b)
}

func logger(h http.Handler) http.Handler {
	fmt.Println("logger")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requested %s", r.RemoteAddr, r.URL)
		hello(w, r)
		//a:=1
		//return a
	})
}

type headerSetter struct {
	key, val string
	handler  http.Handler
}

func (hs headerSetter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ServeHttp2")
	w.Header().Set(hs.key, hs.val)
	hs.handler.ServeHTTP(w, r)
}

// type constructor func(http.Handler) http.Handler

func newHeaderSetter(key, val string) func(http.Handler) http.Handler {
	fmt.Println("newHeaderSetter")
	return func(h http.Handler) http.Handler {   // it satisfies the constructor
		return headerSetter{key, val, h}
	}
}

func main() {
	h := http.NewServeMux()

	h.Handle("/one", numberDumper(1))
	h.Handle("/two", numberDumper(2))
	h.Handle("/three", numberDumper(3))
	h.Handle("/four", numberDumper(4))

	fiveHS := newHeaderSetter("X-FIVE", "the best number")
	h.Handle("/five", fiveHS(numberDumper(5)))

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "That's not a supported number!")
	})

	chain := alice.New(
		newHeaderSetter("X-FOO", "BAR"),
		newHeaderSetter("X-BAZ", "BUZ"),
		logger,
	).Then(h)

	err := http.ListenAndServe(":9999", chain)
	log.Fatal(err)
}