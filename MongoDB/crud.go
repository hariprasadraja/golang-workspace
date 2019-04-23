package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

/* Handling with Struct */
/* It will create a new document Collection*/
func Create(book interface{}, c *mgo.Collection) {
	fmt.Println("Create a book")
	c.Insert(book)

}

func Retrive(book interface{}, c *mgo.Collection) {
	// for Retriving all the elements
	fmt.Println("inside the Retrive method")
	var output interface{}
	// for Retriving only one element
	c.Find(bson.M{}).One(&output)
	fmt.Println(output)
}
func Update(book interface{}, c *mgo.Collection) {
	fmt.Println("Inside the update method")
	c.Update(bson.M{"description": "Building Scallable web apps with Restfull services"},
		bson.M{"$set": bson.M{"description": "It is updated"}})
	c.Find(bson.M{}).One(book)

	fmt.Println(book)
}
func Delete(book interface{}, c *mgo.Collection) {
	fmt.Println("Inside the delete")
	err := c.Remove(bson.M{"isbn": 7894651})
	if err != nil {
		fmt.Printf("Error occurs", err)
		os.Exit(1)
		// we can use panic instead of os.Exit(1)
		fmt.Printf("Collection is deleted")
	}
}
func Retriveall(book interface{}, c *mgo.Collection) {
	var output []interface{}
	fmt.Println("Retrive all")
	c.Find(bson.M{}).All(&output)
	fmt.Println(output)
}
func Updateall(book interface{}, c *mgo.Collection) {
	fmt.Println("update all")
	c.UpdateAll(bson.M{"description": "Building Scallable web apps with Restfull services"},
		bson.M{"$set": bson.M{"description": "It is updated"}})

}
func Deleteall(book interface{}, c *mgo.Collection) {
	fmt.Println("Inside the remove all")
	c.RemoveAll(bson.M{})
	fmt.Println("all docments are removed")
}

func Retriveonbyone(book interface{}, c *mgo.Collection) {
	fmt.Println("Inside the Retriveonbyone")
	var output interface{}
	iter := c.Find(bson.M{}).Iter()
	for iter.Next(&output) {
		fmt.Println("Document ---", output)

	}

}

/*---------------------------------------*/
/* Handling map functions*/

type Book struct {
	Title       string
	Description string
}

func main() {
	session, error := mgo.Dial("localhost:27017")
	defer session.Close()
	if error != nil {
		panic("Error has been occured")
	}

	// to create an Collection

	bookstruct := &Book{Title: "Webdevelopment with go",
		Description: "Building Scallable web apps with Restfull services",
		Author:      "Shiju Varghese",
		ISBN:        7894651}
	collection := session.DB("Library").C("Technology")
	fmt.Println(bookstruct)
	fmt.Println(collection)
	Create(bookstruct, collection)
	Retrive(bookstruct, collection)
	Update(bookstruct, collection)
	//  Delete(bookstruct,collection)
	Retriveall(bookstruct, collection)
	Updateall(bookstruct, collection)
	//    Deleteall(bookstruct,collection)
	Retriveonbyone(bookstruct, collection)
	bookmap := map[string]interface{}{"title": "Web development with go",
		"description": "Building Scallable web apps with Restfull services",
		"author":      "Shiju Vargese",
		"isbn":        7894651}
	fmt.Println("Handling with maps")
	Create(bookmap, collection)
	Retrive(bookmap, collection)
	Update(bookmap, collection)
	//    Delete(bookmap,collection)
	Retriveall(bookmap, collection)
	Updateall(bookmap, collection)
	//  Deleteall(bookmap,collection)
	Retriveonbyone(bookmap, collection)

}
