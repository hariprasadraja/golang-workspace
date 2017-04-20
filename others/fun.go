package main

import (
"fmt"
	"log"
)

type ider interface {
	ID() string
}

type Field struct {
	Name string
}

// Implements ider interface
func (f Field) ID() string {
	return f.Name
}

// uses slice of ider
func inSlice(idSlice []ider, s string) bool {
	log.Println(idSlice)
     x := idSlice[0].ID()
	 if x == s {
		 return true
	 }

	//for _, v := range idSlice {
	//	if s == v.ID() {
	//		return true
	//	}


	return false
}

func main() {
	 var data = make([]ider,5)
       //var fields = Field{"Field1"}
       //data = []ider(fields)

	fmt.Println(inSlice(data, "Field1"))
}
