package main

import (
	"log"
	"zenpepper.com/zenpepper/server/utils"
)

func main() {

	res := utils.SHAEncode("123456789")
	log.Println(res)
}
