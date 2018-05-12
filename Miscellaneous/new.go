package main

import (
	"log"
)

func f() (string, int) {
	return "a", 2
}

func main() {
	log.Println(f())
}
