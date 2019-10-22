package main

import (
	"fmt"
	"log"
)

func hammingDistance(x int, y int) int {
	z := x ^ y
	log.Print("z: \n", z)

	n := 0
	for z != 0 {
		z &= z - 1
		fmt.Printf("Z: %v\n", z)
		n++
	}
	return n
}

func main() {
	fmt.Print(hammingDistance(1, 4))
}
