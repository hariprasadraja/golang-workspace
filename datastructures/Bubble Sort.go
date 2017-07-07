package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	var runes []rune = []rune{'a', 'G', 'c', 'F', 'z', 'D', 'h'}
	fmt.Println("Unsorted:", string(runes))
	bubbleSort(runes)
	fmt.Println("Sorted:", string(runes))
	debug.PrintStack()

}

func bubbleSort(numbers []rune) {

	var n = len(numbers)
	for i := 0; i < n; i++ {
		swap(numbers)
	}
	debug.PrintStack()

}

func swap(numbers []rune) {
	var N = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1

	for secondIndex < N {
		var firstNumber rune = numbers[firstIndex]
		var secondNumber rune = numbers[secondIndex]

		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
		}

		firstIndex++
		secondIndex++
	}
}
