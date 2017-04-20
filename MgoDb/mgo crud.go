package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
)




func main() {
     // Data which is going to store in data base
	direction := map[string]int{
		"North": 3711,
		"East":   2138,
		"South": 1908,
		"West": 912,
	}

	var Direction interface{}

	DB := connectDatabase()

	Cn := DB.C("testing")

	DB.C("testing").Insert(direction)
	Cn.Find(bson.M{"North":3711}).One(Direction)
	log.Println(Direction)

	Cn.Find(bson.M{"North":3711}).All(Direction)

	// Prints list of collection found in DB
	log.Println(Cn.Find(bson.M{"North":3711}).Count())


	travel := map[string]interface{}{
		"_id":bson.NewObjectId(), // Create a new object id for document in mgodb
		"North": 100,
		"East": 200,
	}

	DB.C("testing").Insert(travel)
	//DB.C("testing").
}


func connectDatabase() *mgo.Database{
	session,err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Print(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session.DB("learn")
}
