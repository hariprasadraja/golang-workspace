package main

import (
	"fmt"
	"os"
	"strconv"
)

func rangeSlice(start, stop int) []int {
	if start > stop {
		panic("Slice ends before it started")
	}
	xs := make([]int, stop-start)
	for i := 0; i < len(xs); i++ {
		xs[i] = i + 1 + start
	}
	return xs
}

func permutation(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Add 1 numeric value ex. go run permutation.go 2")
		os.Exit(1)
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	noOfPerms := permutation(rangeSlice(0, num))
	fmt.Println("Number of Permutation:", len(noOfPerms))
	for i := 0; i < len(noOfPerms); i++ {
		fmt.Println(noOfPerms[i])
	}
}
