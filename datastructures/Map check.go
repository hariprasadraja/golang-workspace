package datastructure

import (
	"log"
)

func main() {

	type Check struct {
		key   interface{}
		value interface{}
	}

	c := make(map[interface{}]interface{})
	check := Check{"hi", 123}

	c[check] = check

	log.Println(c[check])
}
