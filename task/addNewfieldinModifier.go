package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"zenpepper.com/zenpepper/server/dbcon"
	"zenpepper.com/zenpepper/server/models"
)

func main() {
	db := connectDatabase("roo")
	menuItem := []models.MenuItem{}

	db.C(dbcon.MenuItem).Find(bson.M{"name": "PEPPERONI PIZZA"}).All(&menuItem)
	for _, item := range menuItem {
		for _, config := range item.CustomConfigs {
			log.Println("severing Size", config.Name)
			config.OptionalModifiers[1] = config.OptionalModifiers[0]
			for _, mg := range config.OptionalModifiers {
				log.Println("ModifierGroup", mg.Name)
				for _, modifier := range mg.Modifiers {
					log.Println("modifier", modifier.Name)
				}
			}
		}
	}
}

func connectDatabase(dbName string) *mgo.Database {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Print(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session.DB(dbName)
}
