package main

import (
	"compress/gzip"
	"encoding/json"
	"net/http"
	//"log"
	"log"
)

func main() {

	http.HandleFunc("/", ResponseHandler)
	http.ListenAndServe(":9999", nil)

}

func ResponseHandler(w http.ResponseWriter, r *http.Request) {
	//var render interface{}
	msg := "This si the response"
	res := make(map[string]interface{})
	res["message"] = msg
	resByte, _ := json.MarshalIndent(res, "", "	")
	//render = msg
	//log.Println(render)
	log.Println(resByte)
	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")
	w.WriteHeader(http.StatusOK)
	// Gzip data
	gz := gzip.NewWriter(w)
	//gz.Write(render.([]byte))
	gz.Write(resByte)
	gz.Close()

}

func JsonGetHandler(w http.ResponseWriter, r *http.Request) {
	// create header

}

func RenderJSON(w http.ResponseWriter, status int, res interface{}) {
	resByte, _ := json.MarshalIndent(res, "", "	")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(resByte)
}
