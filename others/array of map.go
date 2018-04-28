package main

import (
	"fmt"
	"log"
)

func main() {

	modifier := make(map[string][]map[string]interface{})

	maps := map[string]interface{}{
		"name": 20,
	}
	modifier["data"][0] = maps

	//something := []string{"coins", "notes", "gold?"}
	//
	//thisMap := make(map[string][]map[string]int)
	//
	//for _, v := range something {
	//	fmt.Println(v)
	//	aMap := map[string]int{
	//		"random": 12,
	//	}
	//
	//	thisMap[v] = append(thisMap[v], aMap)
	//	fmt.Println(thisMap[v])
	//}
	//log.Println("outside")
	//log.Println(thisMap["coins"][0]["random"])

}
