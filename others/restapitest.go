package main

import (
	"fmt"
	"net/http"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
  
	http.ListenAndServe(":8080",ServeHTTP)
}
