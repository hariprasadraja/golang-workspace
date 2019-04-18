// Rotating array left - using reverse algorithm
// Complexity O(n)
// Space complexity O(1)
// Source: https://raw.githubusercontent.com/astromahi/hackerrank/master/array_rotation/rotate_array_left_04.go
package main

import (
	"log"
)

func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79}
	rotate := 3

	rotate = rotate % len(primes)
	if rotate == len(primes) || rotate == 0 {
		log.Printf("Error: invalid rotation indexes", rotate)
		return
	}

	// Right Rotation
	primes = append(primes[len(primes)-rotate:], primes[:len(primes)-rotate]...)

	// TEMP: Snippet for debugging. remove it before commit
	log.Printf("\n right-rotation :: %+v \n\n", primes)

	// Left Rotation       3:                     0:3
	primes = append(primes[rotate:], primes[0:rotate]...)

	// TEMP: Snippet for debugging. remove it before commit
	log.Printf("\n left-rotation :: %+v \n\n", primes)
}
