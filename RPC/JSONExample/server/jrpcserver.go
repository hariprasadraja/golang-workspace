package main

import (
	"golang-workspace/RPC/JSONExample/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
	arith := new(service.Arith)
	err := s.RegisterService(arith, "Arith")
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.Handle("/rpc", s)
	err = http.ListenAndServe(":1234", r)
	if err != nil {
		log.Fatal(err)
	}
}
