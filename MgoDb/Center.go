package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	DB := connectDatabase()
	myLong, myLat := 77.587073, 12.958794
	pipe := DB.C("store").Pipe([]bson.M{
		{"$project": bson.M{"location": bson.M{"type": bson.M{"$literal": "Point"}, "coordinates": []string{"$longitude", "$latitude"}}}},
		{"$match": bson.M{"location.coordinates": bson.M{"$geoWithin": bson.M{"$centerSphere": []interface{}{[]float64{myLong, myLat}, 10 / 6378.1}}}}},
	})
	resp :=[]bson.M{}
	err := pipe.All(&resp)
	if err != nil{
		panic(err)
	}


	log.Println("check 1:",resp)
	log.Println(len(resp))


	pipe = DB.C("store").Pipe([]bson.M{
		{"$project": bson.M{"location": bson.M{"type": bson.M{"$literal": "Point"}, "coordinates": []string{"$longitude", "$latitude"}}}},
		{"$match": bson.M{"location.coordinates": bson.M{"$geoWithin": bson.M{"$center": []interface{}{[]float64{myLong, myLat},10 / 6378.1}}}}},
	})

	err = pipe.All(&resp)
	if err != nil{
		panic(err)
	}
	log.Println("check 2:",resp)
	log.Println(len(resp))

}

func connectDatabase() *mgo.Database{
      session,err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Print(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session.DB("roo")
}
