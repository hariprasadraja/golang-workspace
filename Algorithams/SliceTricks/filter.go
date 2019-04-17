package main

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79}

	// TEMP: Snippet for debugging. remove it before commit
	spew.Print("\n" + strings.Repeat("-", 15) + "	Debug Begin: primes" + strings.Repeat("-", 15) + "\n")
	spew.Dump(primes)
	spew.Print("\n" + strings.Repeat("-", 15) + "	Debug End: primes	" + strings.Repeat("-", 15) + "\n")

	b := primes[:0]

	// TEMP: Snippet for debugging. remove it before commit
	spew.Print("\n" + strings.Repeat("-", 15) + "	Debug Begin: b" + strings.Repeat("-", 15) + "\n")
	spew.Dump(b)
	spew.Print("\n" + strings.Repeat("-", 15) + "	Debug End: b	" + strings.Repeat("-", 15) + "\n")

	for _, val := range primes {
		if ok, validValue := filter(val); ok {
			b = append(b, validValue)
		}
	}

	// TEMP: Snippet for debugging. remove it before commit
	spew.Print("\n" + strings.Repeat("-", 15) + "	Debug Begin: b" + strings.Repeat("-", 15) + "\n")
	spew.Dump(b)
	spew.Print("\n" + strings.Repeat("-", 15) + "	Debug End: b	" + strings.Repeat("-", 15) + "\n")

	// TEMP: Snippet for debugging. remove it before commit
	spew.Print("\n" + strings.Repeat("-", 15) + "	Debug Begin: primes check" + strings.Repeat("-", 15) + "\n")
	spew.Dump(primes)
	spew.Print("\n" + strings.Repeat("-", 15) + "	Debug End: primes check	" + strings.Repeat("-", 15) + "\n")

	// TEMP: Snippet for debugging. remove it before commit
	spew.Print("\n" + strings.Repeat("-", 15) + "	Debug Begin: b" + strings.Repeat("-", 15) + "\n")
	spew.Dump(b)
	spew.Print("\n" + strings.Repeat("-", 15) + "	Debug End: b	" + strings.Repeat("-", 15) + "\n")

	b[0] = 5

	// TEMP: Snippet for debugging. remove it before commit
	spew.Print("\n" + strings.Repeat("-", 15) + "	Debug Begin: b" + strings.Repeat("-", 15) + "\n")
	spew.Dump(b)
	spew.Print("\n" + strings.Repeat("-", 15) + "	Debug End: b	" + strings.Repeat("-", 15) + "\n")
}

func filter(x int) (bool, int) {
	if x%2 == 0 {
		return true, x
	}

	return false, 0
}

func check() {
	b := primes[:0]
	for i, val := range primes {
		if ok, validValue := filter(val); ok {
			b = append(b, validValue)
		} else {

			// XXX: garbage collect the unwanted elements in the underlyig array
			primes[i] = 0 // nil or the zero value of T
		}
	}
}
