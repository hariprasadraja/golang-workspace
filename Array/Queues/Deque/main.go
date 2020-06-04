// Double Ended Queue

package main

import (
	"fmt"
)

// References:
// https://en.wikipedia.org/wiki/Double-ended_queue

/*
 The Array deque data structure allows for efficient addition and removal at both ends. This structure implements the List interface by using the same circular array technique used to represent an ArrayQueue.
*/

// Complexity O(1 + min(i,n-i)) , where
// i is the position we are going to insert
// n is the lenght of the queue
// in other words, O(insertion + minimum_iteration_required_before_insertion)

type Queue struct {
	ReadIndex  int
	WriteIndex int
	buffer     []interface{}
}

// New initialize and returns a new Queue
func New(size uint) Queue {
	return Queue{
		ReadIndex:  0,
		WriteIndex: 0,
		buffer:     make([]interface{}, size),
	}
}

func (q *Queue) Get(i int) interface{} {
	return q.buffer[(q.ReadIndex+i)%len(q.buffer)]
}

func (q *Queue) Set(i int, x interface{}) interface{} {
	y := q.buffer[(q.ReadIndex+i)%len(q.buffer)]
	q.buffer[(q.ReadIndex+i)%len(q.buffer)] = x
	return y
}

func (q *Queue) Add(i int, x interface{}) {
	if q.WriteIndex == len(q.buffer) {
		fmt.Print("Resizing")
		q.Resize()
	}

	// 2 < 3
	if i < q.WriteIndex/2 {
		q.ReadIndex = ((q.ReadIndex - 1) % len(q.buffer))

		// complexity: O(i)
		for k := 1; k < i; k++ {
			q.buffer[(q.ReadIndex+k)%len(q.buffer)] = q.buffer[(q.ReadIndex+k+1)%len(q.buffer)]
		}
	} else {
		// complexity: O(WriteIndex-i)
		for k := q.WriteIndex; k > i+1; k-- {
			q.buffer[(q.ReadIndex+k)%len(q.buffer)] = q.buffer[(q.ReadIndex+k-1)%len(q.buffer)]
		}
	}

	q.buffer[(q.ReadIndex+i)%len(q.buffer)] = x
	q.WriteIndex += 1
}

func (q *Queue) Remove(i int, x interface{}) interface{} {
	x = q.buffer[(q.ReadIndex+1)%len(q.buffer)]

	if i < q.WriteIndex/2 {
		// q.ReadIndex = (q.ReadIndex - 1) % len(q.buffer)
		for k := i; k >= 1; k-- {
			q.buffer[(q.ReadIndex+k)%len(q.buffer)] = q.buffer[(q.ReadIndex+k+1)%len(q.buffer)]
		}

		q.ReadIndex = ((q.ReadIndex + 1) % len(q.buffer))
	} else {
		for k := 0; k < len(q.buffer)-2; k++ {
			q.buffer[(q.ReadIndex+k)%len(q.buffer)] = q.buffer[(q.ReadIndex+k+1)%len(q.buffer)]
		}
	}

	q.WriteIndex -= 1
	if len(q.buffer) >= 3*q.WriteIndex {
		q.Resize()
	}

	return x
}

func (q *Queue) Resize() {
	// creates a new buffer with double the size of existing buffer
	newbuffer := make([]interface{}, 2*q.WriteIndex)
	for i := 0; i < len(q.buffer); i++ {
		newbuffer[i] = q.buffer[(q.ReadIndex+i)%len(q.buffer)]
	}

	q.buffer = newbuffer
	q.ReadIndex = 0
}

func main() {
	queue := New(6)
	queue.Add(0, "a")
	fmt.Printf("%+v \n", queue)

	queue.Add(1, "b")
	fmt.Printf("%+v \n", queue)

	queue.Add(2, "c")
	fmt.Printf("%+v \n", queue)

	queue.Add(3, "d")
	fmt.Printf("%+v \n", queue)

	queue.Add(2, "e")
	fmt.Printf("%+v \n", queue)

	queue.Add(5, "f")
	fmt.Printf("%+v \n", queue)

	queue.Add(2, "g")
	fmt.Printf("%+v \n", queue)

	queue.Add(3, "3")
	fmt.Printf("%+v \n", queue)
}
