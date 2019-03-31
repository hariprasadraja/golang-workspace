package main

import (
	"log"
)

func main() {
	txt := "Hello World!"
	log.Printf("reversed string: %s", Reverse(txt))
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
