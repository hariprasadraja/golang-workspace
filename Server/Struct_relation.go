package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

const DBName = "RooCheck"
const StoreName = "store"
const ZoneName = "Zone"
const OrderName = "Order"

type DbOperation interface {
	Save() error
	Update(id bson.ObjectId) error
	Remove(id bson.ObjectId) error
	//Archive()
}

type Store struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	DateCreated time.Time     `bson:"dateCreated"`
	DateUpdated time.Time     `bson:"dateUpdated"`
}

type Zone struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	DateCreated time.Time     `bson:"dateCreated"`
	DateUpdated time.Time     `bson:"dateUpdated"`
}

type Order struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	DateCreated time.Time     `bson:"dateCreated"`
	DateUpdated time.Time     `bson:"dateUpdated"`
}

func main() {

	http.HandleFunc("insertStore", InsertStore)
	http.HandleFunc("updateStore", UpdateStore)
	http.HandleFunc("removeStore", RemoveStore)
	http.HandleFunc("insertOrder", InsertOrder)
	http.HandleFunc("updateOrder", UpdateOrder)
	http.HandleFunc("removeOrder", RemoveOrder)
	http.HandleFunc("insertOrder", InsertZone)
	http.HandleFunc("updateOrder", UpdateZone)
	http.HandleFunc("removeOrder", RemoveZone)

	http.ListenAndServe("1234", nil)
}

func InsertStore(w http.ResponseWriter, r *http.Request) {

}

func UpdateStore(w http.ResponseWriter, r *http.Request) {

}

func RemoveStore(w http.ResponseWriter, r *http.Request) {

}

func InsertOrder(w http.ResponseWriter, r *http.Request) {

}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {

}

func RemoveOrder(w http.ResponseWriter, r *http.Request) {

}

func InsertZone(w http.ResponseWriter, r *http.Request) {

}

func UpdateZone(w http.ResponseWriter, r *http.Request) {

}

func RemoveZone(w http.ResponseWriter, r *http.Request) {

}

func connectDatabase() *mgo.Session {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Print(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}
