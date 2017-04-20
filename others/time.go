package main

import (
	"time"
	"fmt"
)

func main() {
	c := time.Tick(1 * time.Minute)
	for now := range c {
		fmt.Printf("%v %s\n", now)
	}

}
