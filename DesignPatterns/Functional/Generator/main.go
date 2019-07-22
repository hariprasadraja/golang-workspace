// This file explains Generator Pattern in golang.
// source: https://github.com/tensor-programming/pattern-tutorial/blob/master/generator/main.go

package main

import "fmt"

func fibGenerator(n int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}

	}()

	return out
}

func main() {
	for x := range fibGenerator(10000000) {
		fmt.Println(x)
	}
}
