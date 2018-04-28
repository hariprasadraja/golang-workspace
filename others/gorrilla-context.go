package main

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
)

type Key string

const GlobalRequestVariable Key = ""

func SetGlobalHandler(w http.ResponseWriter, r *http.Request) {
	context.Set(r, GlobalRequestVariable, "test")

	// will WORK because within request lifetime
	get, ok := context.GetOk(r, GlobalRequestVariable)
	w.Write([]byte(fmt.Sprintf("GetOK : [%v] and get what :[%v] ", ok, get)))

	// global variable still accessible as long the request is still alive
	// and everything is cleared because router from gorilla/mux will
	// automatically call context.Clear(r)
	// and cause GetGlobalHandler to fail
	// WHEN you point the browser url to example.com:8080/get

	// to make it work, you have to pass r down to inside SetGlobalHandler
	// while the request is still alive to InternalGetGlobalHandler
	InternalGetGlobalHandler(w, r)

}

func InternalGetGlobalHandler(w http.ResponseWriter, r *http.Request) {

	// will WORK because still within request lifetime
	get, ok := context.GetOk(r, GlobalRequestVariable)
	w.Write([]byte(fmt.Sprintf("\nInternal GetOK : [%v] and get what :[%v] ", ok, get)))

}

func GetGlobalHandler(w http.ResponseWriter, r *http.Request) {

	// will FAIL because NOT within request lifetime
	get, ok := context.GetOk(r, GlobalRequestVariable)
	w.Write([]byte(fmt.Sprintf("GetOK : [%v] and get what :[%v] ", ok, get)))

}

func main() {
	mx := mux.NewRouter()

	mx.HandleFunc("/", SetGlobalHandler)
	mx.HandleFunc("/get", GetGlobalHandler)

	http.ListenAndServe(":8080", mx)
}
