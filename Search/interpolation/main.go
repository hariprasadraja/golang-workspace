package main

import (
	"fmt"
)

/*
InterpolationSearch searches for the entity in the given sortedData.
if the entity is present, it will return the index of the entity, if not -1 will be returned.
see: https://en.wikipedia.org/wiki/Interpolation_search
Complexity
	Worst: O(N)
	Average: O(log(log(N))  if the elements are uniformly distributed
	Best: O(1)

Example
		fmt.Println(InterpolationSearch(100, []int{1, 2, 9, 20, 31, 45, 63, 70, 100}))
*/
func InterpolationSearch(entity int, sortedData []int) int {
	var (
		minIndex = 0
		minVal   = sortedData[minIndex]

		maxIndex = len(sortedData) - 1
		maxVal   = sortedData[maxIndex]
	)

	for {
		if entity < minVal || entity > maxVal {
			return -1
		}

		// make a guess of the location
		var guess int
		if maxIndex == minIndex {
			guess = maxIndex
		} else {
			size := maxIndex - minIndex
			offset := int(float64(size-1) * (float64(entity-minVal) / float64(maxVal-minVal)))
			guess = minIndex + offset
		}

		// maybe we found it?
		if sortedData[guess] == entity {
			// scan backwards for start of value range
			for guess > 0 && sortedData[guess-1] == entity {
				guess--
			}
			return guess
		}

		// if we guessed to high, guess lower or vice versa
		if sortedData[guess] > entity {
			maxIndex = guess - 1
			maxVal = sortedData[maxIndex]
		} else {
			minIndex = guess + 1
			minVal = sortedData[minIndex]
		}
	}
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(InterpolationSearch(100, items))
}
