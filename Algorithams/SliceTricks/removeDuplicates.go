package main

import (
	"fmt"
	"sort"
)

func main() {
	primes := []int{2, 3, 5, 7, 2, 3, 5, 7, 11, 13, 17, 11, 13, 17, 2, 3, 5, 7, 11, 13, 17}

	sort.Ints(primes)
	j := 0
	for i := 1; i < len(primes); i++ {
		if primes[j] == primes[i] {
			continue
		}

		j++
		primes[j] = primes[i]
	}

	result := primes[:j+1]
	fmt.Println(result)
}
