package con

import (
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

var Db *mgo.Database

func DbConnection() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("Error in dial mongo db")
		os.Exit(1)
	}

	session.SetMode(mgo.Monotonic, true)
	Db = session.DB("simpleGoServer")
}
