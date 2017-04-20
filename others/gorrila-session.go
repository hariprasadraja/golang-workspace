package main

import (
	//"fmt"
	//"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
	"github.com/gorilla/sessions"
)

// For this code to run, you will need this package:
// go get github.com/nu7hatch/gouuid

func main(){

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9999", nil)
}


func get(w http.ResponseWriter,req *http.Request){

}

func foo(w http.ResponseWriter, req *http.Request) {
	//
	//cookie, err := req.Cookie("session-id")
	//if err != nil {
	//	id, err := uuid.NewV4()
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	cookie = &http.Cookie{
	//		Name:  "session-id",
	//		Value: id.String(),
	//		// Secure: true,
	//		HttpOnly: true,
	//	}
	//	http.SetCookie(w, cookie)
	//}
	//fmt.Println(cookie)

	store := sessions.NewCookieStore([]byte("ahdsjffiwuhajdbnf"))
	log.Println("store:" , store)
	store1 :=sessions.NewCookieStore([]byte("anotherkey"))
	log.Println("store1:" , store1)

	data, err := store.Get(req, "session-name")
	log.Println("data:",data)
	if err != nil {
		log.Println("Error in data")
	}
	data1,err := store.Get(req,"session-name1")
	if err != nil {
		log.Println("Error in data1")
	}
	log.Println("data1:",data1)

	 data.Values["Foo"]="hi"
	 data.Values["Another data"] = "Hello World"
	 data.Save(req,w)
	 w.Write([]byte("Session is created"))
}
