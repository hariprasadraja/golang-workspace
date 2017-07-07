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

	var content []interface{}

	customerDetail := customerInfo{"Name": " 1"}
	content = append(content, customerDetail)

	log.Println("info: ", content)

	db := connectDatabase("test")
	//bulk :=db.C("bulk").Bulk()
	//bulk.Insert(&content)

	//db.C("bulk").Insert(customer)
	bulk := db.C("Bulk").Bulk()
	//bulk.Insert(content...)
	//result,err :=bulk.Run()
	//fmt.Printf("%+v",result)
	//fmt.Printf("Error:",err)
	bulk.RemoveAll(content...)
	_, err := bulk.Run()
	if err != nil {
		log.Println(err)
	}
	//
	//collection := db.C("bulk")
	//testBulkInsert(collection)

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

func connectDatabase(dbName string) *mgo.Database {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Print(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session.DB(dbName)
}
