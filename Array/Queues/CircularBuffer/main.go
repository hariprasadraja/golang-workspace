package main

import "fmt"

// https://github.com/zfjagann/golang-ring/blob/master/ring.go

type CircularBuffer struct {
	Max    uint // end in memory
	Head   uint // start of the valid data
	Tail   uint // end of the valid data
	buffer []interface{}
}

// New creates a new circular buffer
func New(size uint) CircularBuffer {
	return CircularBuffer{
		Max:    size,
		Head:   0,
		Tail:   0,
		buffer: make([]interface{}, size),
	}
}

// Write writes the element in the buffer where the tail is pointing
func (b *CircularBuffer) Write(data interface{}) {
	b.buffer[b.Tail] = data
	b.Tail = (b.Tail + 1) % b.Max
}

// POP removes an element in the buffer where head is pointing.
func (b *CircularBuffer) Read() interface{} {
	data := b.buffer[b.Head]
	b.buffer[b.Head] = nil
	b.Head = (b.Head + 1) % b.Max
	return data
}

func GetPageSize() uint {
	return 10
}

func main() {
	buf := New(5)
	buf.Write(1)
	buf.Write(2)
	buf.Write(3)
	buf.Write(4)
	buf.Write(5)
	fmt.Printf("%+v \n", buf)
	buf.Write(6)
	fmt.Printf("%+v \n", buf)
	fmt.Printf("%+v \n", buf)
	fmt.Println(buf.Read())
	fmt.Printf("%+v \n`", buf)

}
