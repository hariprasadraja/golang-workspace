package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	//"net"
	//"crypto/tls"
)

const (
	Application = "application"
	Category    = "category"
	Discount    = "discount"
	MenuItem    = "menuItem"
	Order       = "order"
	ImageData   = "imageData"
	Store       = "store"
	Tax         = "tax"
	User        = "user"
	//Customer         = "customer"
	Transaction      = "transaction"
	Merchant         = "merchant"
	Gateway          = "gateway"
	Token_Merchant   = "tokenMerchant"
	Token_GateWay    = "tokenGateway"
	TokenTransaction = "tokenTransaction"
	CreditCardToken  = "creditCardToken"
	DeliveryZone     = "deliveryZone"
)

var DBObj = DatabaseConnect()

func DatabaseConnect() *mgo.Database {
	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		log.Printf("Error in connecting mongo DB: %s", err.Error())
	}

	session.SetMode(mgo.Monotonic, true)

	return session.DB("deletestore")
}

func main() {
	ids := []bson.ObjectId{}
	StoreIDs := []struct {
		ID bson.ObjectId `bson:"_id,omitempty"`
	}{}
	log.Println("Delete started")
	applicationID := bson.ObjectIdHex("587deae1676161b90a33b4d2")
	DBObj.C(Store).Find(bson.M{"applicationID": applicationID}).All(&StoreIDs)

	for _, StoreID := range StoreIDs {
		ids = append(ids, StoreID.ID)
		log.Println(ids)
	}
	id := bson.M{"$in": ids}

	collections := []string{Category, Discount, MenuItem, Order, ImageData, Store, Tax, User, Transaction, Merchant, Gateway, Token_Merchant, Token_GateWay, TokenTransaction, CreditCardToken, DeliveryZone}
	for _, v := range collections {
		info, err := DBObj.C(v).RemoveAll(bson.M{"internalStore": id})
		log.Println(v)
		log.Printf("%+v", info)
		if err != nil {
			log.Println(v, ":", err)
		}
	}
	err := DBObj.C(Store).RemoveId(id)
	if err != nil {
		log.Println("store", err)
	}
	err = DBObj.C(Application).RemoveId(applicationID)
	if err != nil {
		log.Println("Application", err)
	}

	log.Println("Delete Finished")
}
