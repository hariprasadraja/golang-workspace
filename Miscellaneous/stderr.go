package main

import (
	"log"
	"os"
)

func main() {
	l := log.New(os.Stderr, "", 1)
	l.Println("log message")
}
