package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	//"log"
)

func main() {
	//handler := func(w http.ResponseWriter, r *http.Request) {
	//	io.WriteString(w, "<html><body>Hello World!</body></html>")
	//}

	req := httptest.NewRequest("GET", "http://localhost:8085/accounts", nil)

	w := httptest.NewRecorder()
	//handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

}
