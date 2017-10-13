package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"lingapos/server/db"
	//"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/bson"
)

var DB = connectDatabase()

type store struct {
	Id      bson.ObjectId `bson:"_id"`
	Account bson.ObjectId `bson:"ownedBy"`
}

func main() {

	stores := []store{}

	res := make(map[string]string)
	res[db.Role] = "store"
	res[db.ReceiptTemplate] = "storeId"
	res[db.StoreJSON] = "store"
	res[db.Reason] = "store"
	res[db.Tax] = "store"
	res[db.Merchant] = "store"
	res[db.PaymentSignature] = "store"
	res[db.CashDrop] = "store"
	res[db.CashierOut] = "store"
	res[db.TillManagement] = "store"
	res[db.EmployeeClockIn] = "store"
	res[db.OverTime] = "storeId"
	res[db.GiveX] = "store"
	res[db.EMVSettings] = "store"
	res[db.Tip] = "store"
	res[db.PaidInOut] = "store"
	res[db.PaymentLog] = "store"
	res[db.PrinterConfig] = "store"
	res[db.Network] = "store"

	DB.C(db.Store).Find(nil).Select(bson.M{"_id": true, "ownedBy": true}).All(&stores)
	//log.Println(stores)

	for _, val := range stores {
		for Db_name, name := range res {
			query := bson.M{name: val.Id}
			err := DB.C(Db_name).Update(query, bson.M{"$set": bson.M{"account": val.Account}})
			if err != nil {
				log.Println(Db_name)
				log.Println("Store Id", val.Id)
				log.Println("Account Id", val.Account)
				log.Println(err)
			}
		}
	}
}
