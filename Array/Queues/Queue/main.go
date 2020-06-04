package main

// References
// 1. http://opendatastructures.org/

import "fmt"

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

func (q *Queue) Add(x interface{}) {
	if q.WriteIndex+1 > len(q.buffer) {
		q.Resize()
	}

	q.buffer[(q.ReadIndex+q.WriteIndex)%len(q.buffer)] = x
	q.WriteIndex += 1
}

func (q *Queue) Remove() interface{} {
	x := q.buffer[q.ReadIndex]
	q.ReadIndex = (q.ReadIndex + 1) % len(q.buffer)
	q.WriteIndex = q.WriteIndex - 1

	// Resize operation is very costly while removing.
	if len(q.buffer) >= 3*q.WriteIndex {
		q.Resize()
	}

	fmt.Print(q.ReadIndex)
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
	queue.Add("a")
	queue.Add("b")
	queue.Add("c")
	queue.Add("d")
	queue.Add("e")
	queue.Add("f")
	queue.Add("g")

	fmt.Printf("%+v", queue)

	queue.Remove()
	queue.Remove()

	fmt.Printf("%+v", queue)
}