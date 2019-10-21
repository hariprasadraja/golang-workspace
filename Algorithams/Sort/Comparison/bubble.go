package comparison

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
