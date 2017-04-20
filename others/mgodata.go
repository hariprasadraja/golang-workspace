package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	session,err := mgo.Dial("localhost:27017")
	data := make(map[string]interface{})

	if err!=nil{
		log.Println("connection failed")
		panic(err)
	}
	//db.getCollection('store').find({
	//	"applicationID":ObjectId("589c6750f0de3a93b895e89e"),
	//		"timings":{$elemMatch:{
	//	"day":2,
	//	"status":true,
	//	"from":{$gte:300},
	//	"to":{$lte:1500}
	//	}
	//	}
	//}).count()

	//filter := make(map[string]interface{})
	//filter["a"]


	session.DB("roo").C("store").Find(bson.M{
		"applicationID":bson.ObjectIdHex("589c6750f0de3a93b895e89e"),
		"timings":bson.M{"$elemMatch":bson.M{"day":2,
				"status":true,
				"from":bson.M{"$gte":300},
			     "to":bson.M{"$lte":1500}, }}}).One(data)
	log.Println(data)


}
