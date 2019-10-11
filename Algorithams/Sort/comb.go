package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := generateSlice(10)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	CombSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

/*
CombSort sorts the given array using the comb sort algoritham
Complexity:
	Worst n^2
	Best n log n
	Average n^2 / 2^P, where P is the number of increments
Reference: https://en.wikipedia.org/wiki/Comb_sort
*/
func CombSort(items []int) {
	var (
		n       = len(items) // total elements in the array
		gap     = len(items) // initilize the gap size
		shrink  = 1.3        // shrink factor (recomended)
		swapped = true
	)

	for swapped {
		swapped = false
		gap = int(float64(gap) / shrink)
		if gap < 1 {
			gap = 1
		}

		for i := 0; i+gap < n; i++ {
			if items[i] > items[i+gap] {
				items[i+gap], items[i] = items[i], items[i+gap]
				swapped = true
			}
		}
	}
}

/* Debug Log

--- Unsorted ---

 [-596 -208 307 -248 236 83 -777 551 459 -97]

===	Debug BEGIN: CombSort	===


 gap= 7


 i=0,gap=7, 7 < 10
 NO SWAP: items[0]: -596 < items[7]: 551
 items= [-596 -208 307 -248 236 83 -777 551 459 -97]


 i=1,gap=7, 8 < 10
 NO SWAP: items[1]: -208 < items[8]: 459
 items= [-596 -208 307 -248 236 83 -777 551 459 -97]


 i=2,gap=7, 9 < 10
 SWAP: items[2]: 307 > items[9]: -97
 items= [-596 -208 -97 -248 236 83 -777 551 459 307]


 gap= 5


 i=0,gap=5, 5 < 10
 NO SWAP: items[0]: -596 < items[5]: 83
 items= [-596 -208 -97 -248 236 83 -777 551 459 307]


 i=1,gap=5, 6 < 10
 SWAP: items[1]: -208 > items[6]: -777
 items= [-596 -777 -97 -248 236 83 -208 551 459 307]


 i=2,gap=5, 7 < 10
 NO SWAP: items[2]: -97 < items[7]: 551
 items= [-596 -777 -97 -248 236 83 -208 551 459 307]


 i=3,gap=5, 8 < 10
 NO SWAP: items[3]: -248 < items[8]: 459
 items= [-596 -777 -97 -248 236 83 -208 551 459 307]


 i=4,gap=5, 9 < 10
 NO SWAP: items[4]: 236 < items[9]: 307
 items= [-596 -777 -97 -248 236 83 -208 551 459 307]


 gap= 3


 i=0,gap=3, 3 < 10
 NO SWAP: items[0]: -596 < items[3]: -248
 items= [-596 -777 -97 -248 236 83 -208 551 459 307]


 i=1,gap=3, 4 < 10
 NO SWAP: items[1]: -777 < items[4]: 236
 items= [-596 -777 -97 -248 236 83 -208 551 459 307]


 i=2,gap=3, 5 < 10
 NO SWAP: items[2]: -97 < items[5]: 83
 items= [-596 -777 -97 -248 236 83 -208 551 459 307]


 i=3,gap=3, 6 < 10
 NO SWAP: items[3]: -248 < items[6]: -208
 items= [-596 -777 -97 -248 236 83 -208 551 459 307]


 i=4,gap=3, 7 < 10
 NO SWAP: items[4]: 236 < items[7]: 551
 items= [-596 -777 -97 -248 236 83 -208 551 459 307]


 i=5,gap=3, 8 < 10
 NO SWAP: items[5]: 83 < items[8]: 459
 items= [-596 -777 -97 -248 236 83 -208 551 459 307]


 i=6,gap=3, 9 < 10
 NO SWAP: items[6]: -208 < items[9]: 307
 items= [-596 -777 -97 -248 236 83 -208 551 459 307]


===	Debug END: CombSort	===


--- Sorted ---

 [-596 -777 -97 -248 236 83 -208 551 459 307]
*/
