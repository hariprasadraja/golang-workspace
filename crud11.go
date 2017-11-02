package main

// Crud operations with mongodb,restapi
import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	// "log"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Ename string        `json:"ename" bson: "ename"`
	Eid   string        `json:"eid" bson: "eid"`
}

func Insert(w http.ResponseWriter, r *http.Request) {
	session, _ := mgo.Dial("localhost:27017")
	person := Person{}

	reqBytes, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBytes, &person)

	err := session.DB("employee").C("emp").Insert(&person)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(person)
	fmt.Println("inserted")
}

func remove(res http.ResponseWriter, req *http.Request) {
	session, _ := mgo.Dial("localhost:27017")
	c := session.DB("employee").C("emp")

	_, err := c.RemoveAll(bson.M{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("removed")
	}
}

func Update(res http.ResponseWriter, r *http.Request) {
	fmt.Println("entry")
	session, _ := mgo.Dial("localhost:27017")
	p1 := Person{}

	// params := context.Get(r, "params").(httprouter.Params)
	// id := params.ByName("5833e1a738eed2cc102e2f02")

	err := session.DB("employee").C("emp").Find(bson.M{"_id": bson.ObjectIdHex("5833e1a738eed2cc102e2f02")}).One(&p1)
	if err != nil {
		fmt.Println("id not fount", err.Error())
	}

	bytes, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(bytes, &p1)

	err = session.DB("employee").C("emp").UpdateId(p1.ID, &p1)
	if err != nil {
		fmt.Println("cout't not update", err.Error())
	}
	// err := c.Update(bson.M{"Eid": "A"}, bson.M{"$set": bson.M{"Ename": "avi"}})
	// if err != nil {
	// 	fmt.Printf("err", err)
	// } else {
	fmt.Println("data's", p1)
	fmt.Printf("updated\n")
}

func main() {

	http.HandleFunc("/remove", remove)

	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.ListenAndServe(":8080", nil)

}
