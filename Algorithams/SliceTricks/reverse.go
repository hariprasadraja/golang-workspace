package main

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79}

	for left, right := 0, len(primes)-1; left < right; left, right = left+1, right-1 {
		primes[left], primes[right] = primes[right], primes[left]
	}

	// TEMP: Snippet for debugging. remove it before commit
	spew.Print("\n" + strings.Repeat("-", 15) + "	Debug Begin: primes" + strings.Repeat("-", 15) + "\n")
	spew.Dump(primes)
	spew.Print("\n" + strings.Repeat("-", 15) + "	Debug End: primes	" + strings.Repeat("-", 15) + "\n")
}
