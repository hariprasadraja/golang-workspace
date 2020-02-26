package main

import "fmt"

func main() {
	num := 5
	len := 2
	for i, draw := range spiral(num) {
		fmt.Printf("%*d ", len, draw)
		if i%num == num-1 {
			fmt.Println("")
		}
	}

	num = 5
	len = 2
	for i, draw := range zigzag(num) {
		fmt.Printf("%*d ", len, draw)
		if i%num == num-1 {
			fmt.Println("")
		}
	}
}
