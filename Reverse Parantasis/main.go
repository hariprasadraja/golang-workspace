package main

import (
	"log"
	"strings"
)

func main() {
	//a := "hi(bar)hi"
	a := "foo(bar(baz))blim"
	open := strings.Count(a, "(")
	open += strings.Count(a, ")")
	if open%2 != 0 {
		log.Println("string can't be parsed")
	}

}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
