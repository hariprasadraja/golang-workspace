package main

import (
	"log"
)

// TotalNodes is the total number of nodes in the graph.
var TotalNodes int

// Graph is represented as Adjacent List
var Graph map[int][]int

// Visited marks  the visited nodes in the graph
var Visited []bool

func DepthFirstSearch(at int) {
	if Visited[at] {
		return
	}

	Visited[at] = true
	log.Println("Visited: ", at)
	neighbours := Graph[at]
	for _, n := range neighbours {
		DepthFirstSearch(n)
	}
}

func main() {
	TotalNodes = 10
	Graph = map[int][]int{
		0: {1, 2, 3},
		1: {4, 5, 3},
		2: {6, 7, 3},
		3: {7, 8, 9},
		4: {5, 2, 3},
		5: {6, 9, 8},
		6: {8, 9, 0},
		7: {2, 4, 5},
		8: {3, 5, 7},
		9: {6, 7, 1},
	}

	Visited = make([]bool, len(Graph))
	DepthFirstSearch(0)
}
