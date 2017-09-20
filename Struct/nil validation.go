package main

import "log"

type Check struct {
	Name string
}

func main() {
	pointer := &Check{}
	variable := Check{}

	if variable == nil {
		log.Print("variable is nil")
	}

	if pointer == nil {
		log.Print("pointer is nil")
	}

}
