package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

// working and usage of Apply function & change struct  in mgo package
var Db = ConnectDatabase("ApplyCheck")
var Collection = "check"
var Direction = map[string]int{
	"North": 0,
	"East":  0,
	"South": 0,
	"West":  0,
}

func main() {

	//Insert()
	UpdateusingApply()

}

//func ConnectDatabase(dbName string) *mgo.Database{
//	session,err := mgo.Dial("localhost:27017")
//	if err != nil {
//		log.Print(err)
//	}
//	session.SetMode(mgo.Monotonic, true)
//	return session.DB(dbName)
//}

func UpdateusingApply() {
	res := make(map[string]int)
	change := mgo.Change{}
	changeinfo := &mgo.ChangeInfo{}

	//for i:=1 ; i <= 20 ; i++ {
	Direction["North"] = 10
	Direction["East"] = 100
	Direction["South"] = 1000
	Direction["West"] = 100000

	change.Update = Direction
	change.ReturnNew = true
	change.Upsert = true

	query := bson.M{
		"North": 10,
		"East":  100,
	}

	Db.C(Collection).Find(query).One(res)
	changeinfo, err := Db.C(Collection).Find(query).Apply(change, res)
	log.Printf("%+v ", changeinfo)
	if err != nil {
		log.Println("error", err)
	}
	for key, val := range res {
		log.Println(key, ":", val)
	}
	//if err != nil {
	//	log.Println("error",err)
	//}
	//}

}

func Insert() {

	for i := 1; i <= 20; i++ {
		Direction["North"] = i
		Direction["East"] = i * 2
		Direction["South"] = i * 3
		Direction["West"] = i * 4

		err := Db.C(Collection).Insert(Direction)
		if err != nil {
			log.Println("error", err)
		}
	}

}

// Results
// 1. if the info is find it is updating. otherwise retrun empty
// 2. SETTING returnNew is set
//       returning updated value
