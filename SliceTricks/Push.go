package main

import "log"

func main() {
	a := []int{1,2,3,4,5,6,7,8,9}

	// Push front
	log.Printf("Array Before Psuh: %+v", a)
	a = append([]int{a[8]}, a...)
	log.Printf("Array After Push: %+v", a)


	a = []int{1,2,3,4,5,6,7,8,9}

	// Right Rotation.
	log.Printf("Array Before Rotation: %+v",a)
	a = append([]int{a[8]}, a[:len(a)-1]...)
	log.Printf("Array After Rotation: %+v", a)

}