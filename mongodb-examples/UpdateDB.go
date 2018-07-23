package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"gopkg.in/mgo.v2"
	"log"
)

func main()  {
	db := connectDatabase()
	db.C("sale").Update(bson.M{"store":bson.ObjectIdHex("56e1c14208811a14c835b73b")},bson.M{"$set":bson.M{"startDate":time.Now().AddDate(0,0,-1)}})
}

func connectDatabase() *mgo.Database {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Print(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session.DB("Linga")
}


