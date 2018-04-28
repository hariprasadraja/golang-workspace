package main

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
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
