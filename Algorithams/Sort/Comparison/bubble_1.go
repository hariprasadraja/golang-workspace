package comparison

func BubbleSort1(numbers []rune) {

	var n = len(numbers)
	for i := 0; i < n; i++ {
		swap(numbers)
	}
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
