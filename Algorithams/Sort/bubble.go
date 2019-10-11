package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := GenerateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	BubbleSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

// GeneratesSlice generates a slice of size filled with random numbers
func GenerateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		// Random data goes here...
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}

	return slice
}

/*
BubbleSort sorts the given unSortedData into the sorted data. If the data is already sorted
then it will return as it is without modification.
Complexity:
	Worst complexity: O(n^2)
	Average complexity: O(n^2)
	Best complexity: O(n)
	Space complexity: O(1)
Reference: https://en.wikipedia.org/wiki/Bubble_sort
*/
func BubbleSort(unSortedData []int) (sortedData []int) {
	var (
		n      = len(unSortedData)
		sorted = false
	)

	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {

			// sort ascending
			if unSortedData[i] > unSortedData[i+1] {
				unSortedData[i+1], unSortedData[i] = unSortedData[i], unSortedData[i+1]
				swapped = true
			}
		}

		// if no swapping occurs, then the array is already sorted
		if !swapped {
			sorted = true
		}

		n = n - 1
	}

	return unSortedData
}
