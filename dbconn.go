package main

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Student struct {
	Id   int    `bson:"id,o"`
	Name string `bson:"-"`
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic("error in connecting Databases")
	}
	defer session.Close()
	c := session.DB("school").C("student")
	/// FOR INSERTION
	c.Insert(&Student{Id: 1, Name: "nagaraj"}, &Student{Id: 2, Name: "sathya"})
	var result []Student
	// FOR RETRIVING
	c.Find(bson.M{}).All(&result)
	fmt.Println(result)
	//FOR REMOVING
	err = c.Remove(bson.M{"id": 1})
	if err != nil {
		fmt.Printf("remove fail %v\n", err)
		os.Exit(1)
	}
	fmt.Println("----------------------------------------------------------")
	var result1 []Student
	c.Find(bson.M{}).All(&result1)
	fmt.Println(result1)
	fmt.Println("----------------------------------------------------------")
	data := Student{}
	ret := c.Find(nil).Iter()
	for ret.Next(&data) {
		fmt.Printf("\nid %d Name %s\n", data.Id, data.Name)
	}
	err = c.Remove(bson.M{"name": "sathya"})
	if err != nil {
		fmt.Printf("remove fail %v\n", err)
		// os.Exit(1)
	}
	c.UpdateAll(bson.M{"name": "sathya"}, bson.M{"$set": bson.M{"name": "kevin", "id": 2}})
	fmt.Println("----------------------------------------------------------")
	data1 := Student{}
	ret2 := c.Find(nil).Sort("name").Iter()
	for ret2.Next(&data1) {
		fmt.Printf("\nid %d Name %s\n", data1.Id, data1.Name)
	}
	fmt.Println("----------------------------------------------------------")
	counts, _ := c.Count()
	fmt.Println("total data is", counts)
	c.RemoveAll(nil)
	fmt.Println("----------------------------------------------------------")
	fmt.Println("total data is", counts)
	fmt.Println("----------------------------------------------------------")
	data2 := Student{}
	ret3 := c.Find(nil).Sort("name").Iter()
	for ret3.Next(&data2) {
		fmt.Printf("\nid %d Name %s\n", data2.Id, data2.Name)
	}

}
