package main

import (
	"crypto/sha256"
	"encoding/base64"
)

func main() {
	SHAEncode("123456789")
}

func SHAEncode(target string) (output string) {
	targetBytes := []byte(target)

	enc := sha256.New()
	enc.Write(targetBytes)

	output = base64.StdEncoding.EncodeToString(enc.Sum(nil))
	println(output)
	return
}
