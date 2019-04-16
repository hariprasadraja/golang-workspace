package main

import (
	"fmt"
	"log"
)

func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79}

	// safe delete
	copy(primes[i:], primes[i+1:])

	// TEMP: Snippet for debugging. remove it before commit
	log.Printf("\n primes-after copy :: %+v \n\n", primes)

	primes[len(primes)-1] = 0
	primes = primes[:len(primes)-1]

	// // TEMP: Snippet for debugging. remove it before commit
	log.Printf("\n primes :: %+v \n\n", primes)

	// cut
	s1 := append(primes[:1], primes[4:]...)
	fmt.Println("Cutted Slice:")
	fmt.Println(s1, len(s1))

	//delete
	i := 5
	fmt.Print("Delete Index: %v", i)
	s2 := append(primes[:i], primes[i+1:]...)
	fmt.Println("Deleted Slice:")
	fmt.Println(s2, len(s2))

	// Delete without preserving order
	s3 := make([]int, len(primes))
	s3[i] = primes[len(primes)-1]
	s3 = primes[:len(primes)-1]

	// TEMP: Snippet for debugging. remove it before commit
	fmt.Printf("\n s3 :: %+v \n", s3)

	// safe Cut for pointers or struct which has pointers to avoid memeory leak
	// i = 5
	// j := 10
	copy(primes[i:], primes[j:])
	for k, n := (len(primes) - j + i), len(primes); k < n; k++ {
		primes[k] = 0
	}

	primes = primes[:len(primes)-j+i]

	// // TEMP: Snippet for debugging. remove it before commit
	// log.Printf("\n primes :: %+v \n\n", primes)

}
