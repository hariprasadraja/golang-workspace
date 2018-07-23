package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

type customerInfo map[string]interface{}

type Content struct {
	Name     string
	Download int
	Date     time.Time
}

func main() {

	init_CategoryBulk()

}

func init_CategoryBulk() {
	categories := []Content{
		{
			Name: "test modifier",
		},
		{
			Name: "test mod there",
		},
	}

	data := []interface{}{}
	for _, category := range categories {
		data = append(data, category)
	}

	bulk := connectDatabase("Bulk").C("category").Bulk()
	bulk.Insert(data...)
	_, err := bulk.Run()
	log.Print(err)
}

func InsertBulk(data []interface{}) {
	db := connectDatabase("BulkInsert")
	bulk := db.C("bulk").Bulk()
	bulk.Insert(data...)

	_, err := bulk.Run()
	if err != nil {
		log.Print("Error in Insert", err.Error())
	}
}

func testBulkInsert(Coll *mgo.Collection) {
	fmt.Println("Test Bulk Insert into MongoDB")
	bulk := Coll.Bulk()

	var contentArray []interface{}
	contentArray = append(contentArray, &Content{
		Name:     "this-is-good-content",
		Download: 1,
		Date:     time.Date(2016, 4, 7, 0, 0, 0, 0, time.UTC),
	})

	contentArray = append(contentArray, &Content{
		Name:     "this-is-good-content",
		Download: 2,
		Date:     time.Now(),
	})

	contentArray = append(contentArray, &Content{
		Name:     "this-is-good-content",
		Download: 2,
		Date:     time.Now(),
	})

	bulk.Insert(contentArray...)
	_, err := bulk.Run()
	if err != nil {
		panic(err)
	}
}
//
//func connectDatabase(dbName string) *mgo.Database {
//	session, err := mgo.Dial("localhost:27017")
//	if err != nil {
//		log.Print(err)
//	}
//	session.SetMode(mgo.Monotonic, true)
//	return session.DB(dbName)
//}
