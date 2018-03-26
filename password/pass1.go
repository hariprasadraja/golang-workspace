package main

import (
	"log"
	"crypto/sha256"
	"encoding/base64"
)


func SHAEncode(target string) (output string) {
	targetBytes := []byte(target)

	enc := sha256.New()
	enc.Write(targetBytes)

	output = base64.StdEncoding.EncodeToString(enc.Sum(nil))

	return
}




func main() {
	res := SHAEncode("callcenter123")
	log.Println(res)
}
