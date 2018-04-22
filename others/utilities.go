package main

import (
	//"math/rand"
	//"time"
	"log"
	//"reflect"
	"math/rand"
)

func main() {
	//rand.Seed(time.Now().Unix())
	//x := time.Now().UnixNano()
	//log.Println(x)
	//log.Println(rand.Int63())
	//log.Println(reflect.TypeOf(x))
	//log.Println(reflect.TypeOf(rand.Int63()))
	//
	var alphanumRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

	for _,id := range alphanumRunes{
		log.Println(id)
		log.Println(len(alphanumRunes))
		log.Println(rand.Intn(len(alphanumRunes)))
	}
}