package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"context"
	"simpleGoServer/server/con"
	"simpleGoServer/server/controllers"
)


func init() {
	con.DbConnection()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("public/index.html")
	greet := map[string]string{"greetings": "Hello world"}
	t.Execute(w, greet)
}

func loggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		status := w.Header().Get("code")
		w.Header().Del("code")
		log.Printf("[%s] %q %v - %s\n", r.Method, r.URL.String(), t2.Sub(t1), status)
	}

	return http.HandlerFunc(fn)
}

// Put params in context for sharing them between handlers
func wrapHandler(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func main() {
	mux := httprouter.New()

	mux.POST("/user", wrapHandler(loggingHandler(http.HandlerFunc(controllers.SaveUser))))
	mux.GET("/user/:id", wrapHandler(loggingHandler(http.HandlerFunc(controllers.GetUser))))

	mux.ServeFiles("/public/*filepath", http.Dir("public"))
	mux.GET("/", wrapHandler(http.HandlerFunc(indexHandler)))

	log.Print("Listening on: 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Print("Error in listening server: ", err)
	}
}
