package main

import "fmt"

/*
BinarySearch (a.k.a. half-interval search/ logarithmic search / binary chop)
searches for an entity which is present in the given data (sorted data).
If the entity is present, it will retrun true along with it's index otherwise false and
the index will be zero
Complexity
		Worst: O(log n)
		Average: O(log n)
		Best: O(1)
		Space: O(1)
*/
func BinarySearch(entity int, sortedData []int) (bool, int) {
	var (
		minIndex = 0
		maxIndex = len(sortedData) - 1
	)

	for minIndex <= maxIndex {
		median := (minIndex + maxIndex) / 2
		if sortedData[median] < entity {

			// if entity is greater than the median, then neglect the data lesser than the median
			minIndex = median + 1
		} else {

			// if entity is lesser than the median, then neglect the data greater than the median
			maxIndex = median - 1
		}
	}

	// if last index reached || value does not match with entity
	if minIndex == len(sortedData) || sortedData[minIndex] != entity {
		return false, 0
	}

	return true, minIndex
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(BinarySearch(63, items))
}
