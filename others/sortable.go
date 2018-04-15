package main

import (
	"fmt"
	"sort"
)

type SortFuncs struct {
	Length func() int
	Lesser func(int, int) bool
	Swaper func(int, int)
}

// type Sortable struct{/ The SortFuncs }

func (a SortFuncs) Len() int           { return a.len() }
func (a SortFuncs) Less(i, j int) bool { return a.less(i, j) }
func (a SortFuncs) Swap(i, j int)      { a.swap(i, j) }

// func ClosureSort(length func() int, less func(int, int) bool, swap func(int, int)) Sortable {
// 	return Sortable{SortFuncs{length, less, swap}}
// }

func main() {

	x := []int{5, 1, 4, 2, 3}
	sort.Sort(SortFuncs{
		Length: func() int { return len(x) },
		Lesser: func(i, j int) bool { return x[i] < x[j] },
		Swaper: func(i, j int) { x[i], x[j] = x[j], x[i] },
	})

	fmt.Print(x)
}
