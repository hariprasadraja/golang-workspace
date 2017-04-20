package main

import (
	"bytes"
	"log"

	//"fmt"

	//"errors"
)

func main() {

log.Println(bytes.MinRead)

	//z.ToLower(rune(90))
	//name :=bytes.ToTitleSpecial(z,[]byte("Hari Prasad"))
	//fmt.Printf("%s\n",name)



	buffer := bytes.Buffer{}
	buffer.Write([]byte("hi"))
	log.Println(buffer.Cap())
	log.Println(buffer.Bytes())
    var a complex64
	a = 10 + 20
    log.Println(a)
	//errorcheck()
	log.Println("recovered")
}

