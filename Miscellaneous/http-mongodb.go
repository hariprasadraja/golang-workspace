package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func CreateHandler(res http.ResponseWriter, req *http.Request) {
	// Handling with Form Values
	log.Println("brand", req.FormValue("brand"))
	log.Println("pages", req.FormValue("pages"))
	log.Println("cost", req.FormValue("cost"))
	log.Println("design", req.FormValue("design"))
	pages, _ := strconv.Atoi(req.FormValue("pages"))
	cost, _ := strconv.Atoi(req.FormValue("cost"))
	book := &Notebook{
		BrandName: req.FormValue("brand"),
		Pages:     pages,
		Cost:      cost,
		Design:    req.FormValue("design"),
	}
	CreateData(book)
	log.Println("exit - CreateHandler")
	fmt.Fprintln(res, "New book is created")
}
func RetriveHandler(res http.ResponseWriter, req *http.Request) {

	book := RetriveData()
	log.Println(book)
}
func UpdateHandler(res http.ResponseWriter, req *http.Request) {
	text := req.FormValue("toupdate")
	log.Println("text to update", text)
	UpdateData(text)
	log.Println("Data is updated")

}
func RemoveHandler(res http.ResponseWriter, req *http.Request) {
	RemoveData()
	log.Println("Data is removed")
}
func CreateData(book *Notebook) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		os.Exit(1)
	}
	session.DB("http-mongodb").C("Notebook").Insert(book)
	log.Println("book is inserted")
	defer session.Close()
}

func RetriveData() *Notebook {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		os.Exit(2)
	}
	var output Notebook
	session.DB("http-mongodb").C("Notebook").Find(nil).One(&output)
	defer session.Close()
	return &output
}
func UpdateData(b string) {
	log.Println(b)
	session, err := mgo.Dial("localhost:27017")
	defer session.Close()
	if err != nil {
		os.Exit(3)
	}

	session.DB("http-mongodb").C("Notebook").Update(bson.M{"brandname": b}, bson.M{"$set": bson.M{"brandname": "content is updated"}})
	log.Println("Data is updated in the database")
}

func RemoveData() {
	session, err := mgo.Dial("localhost:27017")
	defer session.Close()

	if err != nil {
		os.Exit(4)
	}
	session.DB("http-mongodb").C("Notebook").RemoveAll(bson.M{})
}

type Notebook struct {
	BrandName string `json:"brandname"`
	Pages     int    `json:pages`
	Cost      int    `json:cost`
	Design    string `json:design`
}

func main() {
	http.HandleFunc("/createjs/", JsonCreate)
	http.HandleFunc("/viewjs/", JsonView)
	http.HandleFunc("/updatejs/", JsonUpdate)
	http.HandleFunc("/Deletejs/", JsonDelete)

	http.HandleFunc("/create/", CreateHandler)
	http.HandleFunc("/view/", RetriveHandler)
	http.HandleFunc("/update/", UpdateHandler)
	http.HandleFunc("/Delete/", RemoveHandler)
	http.ListenAndServe(":8080", nil)
}

func JsonCreate(res http.ResponseWriter, req *http.Request) {

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		os.Exit(5)
	}
	var p Notebook
	json.Unmarshal(data, &p)

	fmt.Fprintln(res, p.Cost)
	CreateData(&p)
	log.Println("Completed JsonCreate")

}
func JsonView(res http.ResponseWriter, req *http.Request) {
	data := RetriveData()
	log.Println(data)
	fmt.Fprintf(res, "Json view")
	log.Println("Completed JsonView")
}
