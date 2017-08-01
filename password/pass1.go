package main

import (
	"log"
	"zenpepper.com/zenpepper/server/utils"
)

func main() {

	res := utils.SHAEncode("welcome123")
	log.Println(res)
}
