package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

const DBName = "RooCheck"
const StoreName = "store"
const ZoneName = "Zone"
const OrderName = "Order"

type DB struct {
	collection *mgo.Collection
	child      interface {
		validate()
	}
}

type Store struct {
	//DB 					`bson:"-"`
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"storeName"`
	DateCreated time.Time     `bson:"dateCreated"`
	DateUpdated time.Time     `bson:"dateUpdated"`
}

type DeliveryZone struct {
	//DB					 `bson:"-"`
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"zoneName"`
	DateCreated time.Time     `bson:"datecreated"`
	DateUpdated time.Time     `bson:"dateUpdated"`
}

type Order struct {
	//DB					 `bson:"-"`
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"OrderName"`
	DateCreated time.Time     `bson:"datecreated"`
	DateUpdated time.Time     `bson:"dateUpdated"`
}

func (store *Store) validate() {
	log.Println("store validated")
}

func (dz *DeliveryZone) validate() {
	log.Println("delivery zone valildated")
}

func (dz *Order) validate() {
	log.Println("Order valildated")
}

func (db DB) Save() (bson.ObjectId, error) {
	var result map[string]bson.ObjectId
	db.child.validate()
	err := db.collection.Insert(db.child)
	if err != nil {
		return result["_id"], err
	}
	err = db.collection.Find(db.child).Select(bson.M{"_id": 1}).One(&result)
	return result["_id"], err
}

func (db DB) Update(id bson.ObjectId) error {
	db.child.validate()
	err := db.collection.UpdateId(id, db.child)
	return err
}

func (db DB) Remove(id bson.ObjectId) error {
	err := db.collection.RemoveId(id)
	return err
}

func main() {

	session := connectDatabase()
	store := &Store{}
	store.Name = "My Store"
	//store.Id = bson.ObjectIdHex("595de2d20000000000000000")
	db := &DB{}
	db.collection = session.DB(DBName).C(StoreName)

	store.DateCreated = time.Now()
	db.child = store
	id, err := db.Save()
	if err != nil {
		log.Println("Error in Save Store", err.Error())
		return
	}
	store.Id = id

	if store.Id.Hex() != "" {
		store.DateUpdated = time.Now()
		store.Name = "My Updated Store"
		db.child = store
		err := db.Update(store.Id)
		if err != nil {
			log.Println(" Error in Update Store", err.Error())
		}
		//render json
		//return
	}

	err = db.Remove(store.Id)
	if err != nil {
		log.Println("Error in removing store", err.Error())
	}

	zone := &DeliveryZone{}
	zone.Name = "My Zone"
	//store.Id = bson.ObjectIdHex("595de2d20000000000000000")
	db = &DB{}
	db.collection = session.DB(DBName).C(ZoneName)

	zone.DateCreated = time.Now()
	db.child = zone
	id, err = db.Save()
	if err != nil {
		log.Println("Error in Save Zone", err.Error())
		return
	}

	zone.Id = id

	if zone.Id.Hex() != "" {
		zone.Name = "Zone Updated"
		db.child = zone
		err := db.Update(zone.Id)
		if err != nil {
			log.Println(" Error in Update Zone", err.Error())
		}
		//render json
		//return
	}

	err = db.Remove(zone.Id)
	if err != nil {
		log.Println("Error in removing store", err.Error())
	}

	order := &Order{}
	order.Name = "My Order"
	//store.Id = bson.ObjectIdHex("595de2d20000000000000000")
	db = &DB{}
	db.collection = session.DB(DBName).C(OrderName)

	store.DateCreated = time.Now()
	db.child = order
	id, err = db.Save()
	if err != nil {
		log.Println("Error in Save Order", err.Error())
		return
	}
	order.Id = id

	if order.Id.Hex() != "" {
		order.DateUpdated = time.Now()
		order.Name = "Order Updated"
		db.child = order
		err = db.Update(order.Id)
		if err != nil {
			log.Println(" Error in Update Order", err.Error())
		}
		//render json
		//return
	}
	err = db.Remove(order.Id)
	if err != nil {
		log.Println("Error in removing store", err.Error())
	}

}

func connectDatabase() *mgo.Session {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Print(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}
