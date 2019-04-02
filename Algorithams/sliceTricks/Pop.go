package main

import "log"

func main() {

	// POP front
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var x int
	x, a = a[0], a[1:]

	log.Printf("Array: %+v", a)
	log.Printf("first element: %+v", x)

	// Pop Back
	x, a = a[len(a)-1], a[:len(a)-1]
	log.Printf("Array: %+v", a)
	log.Printf("Last Element: %+v", x)
}
