package main

import (
	"bufio"
	"log"
	"bytes"
	"io"
)

func main() {

	var z bool
	buffer := make([]byte, 100)
	buffer = []byte(" Hellow world")
	_, token, err := bufio.ScanBytes(buffer, z)
	// retruns each byte as a token

	log.Println(token)
	// it is a array which contains array of bytes of our given string
	if err != nil {
		panic(err)
	}

	wirter :=new(io.Writer)
	const input= "1234 5678 1234567 9012345 67890"
	Reader := bufio.NewReader(bytes.NewBufferString(input))
	Writer := bufio.NewWriter(*wirter)
	ReaderWriter :=bufio.NewReadWriter(Reader,Writer)
	log.Println(ReaderWriter.Read(buffer))
    log.Println(ReaderWriter.Peek(10))
    log.Println(ReaderWriter.WriteTo(Writer))
	log.Println(Writer.Available())


}
