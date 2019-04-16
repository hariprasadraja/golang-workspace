package main

import "log"

func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79}

	// safe delete
	i := 5 // delete index 5 which has value 13
	copy(primes[i:], primes[i+1:])
	primes[len(primes)-1], primes = 0, primes[:len(primes)-1]

	// TEMP: Snippet for debugging. remove it before commit
	log.Printf("\n primes :: %+v \n\n", primes)

	// delete without order
	primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79}
	primes[i], primes[len(primes)-1], primes = primes[len(primes)-1], 0, primes[:len(primes)-1]

	// TEMP: Snippet for debugging. remove it before commit
	log.Printf("\n primes :: %+v \n\n", primes)

}
