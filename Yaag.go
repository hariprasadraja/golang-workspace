package main

import (
	"github.com/betacraft/yaag/yaag"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
yaag.Init(&yaag.Config{On: true, DocTitle: "Core", DocPath: "apidoc.html", BaseUrls : map[string]string{"Production":"","Staging":""} })
http.HandleFunc("/", middleware.HandleFunc(handler))
http.ListenAndServe(":8080", nil)
}
