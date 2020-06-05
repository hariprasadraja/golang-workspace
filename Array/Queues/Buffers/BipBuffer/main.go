package main

import (
	"fmt"
	"os"
)

// This file in put on HOLD. need to find ways to use mmap

/*
I have taken this from https://www.codeproject.com/script/Articles/ViewDownloads.aspx?aid=3479 and credit goes to Simon Cooke for his beautiful implementation of BipBuffer

Also considered https://github.com/willemt/bipbuffer for implementing this


*/

type BipBuffer struct {
	buffer []interface{}
	ixa    int
	sza    int

	ixb int
	szb int

	buflen int

	ixResrv int
	szResrv int
}

func New() BipBuffer {
	return BipBuffer{}
}

func (b *BipBuffer) AllocateBuffer(bufferSize int) bool {
	// default buffer size, if no size is given
	if bufferSize == 0 {
		bufferSize = 4096
	}

	// Free the buffer
	if b.buffer != nil {
		b.buffer = make([]interface{}, bufferSize)
	}

	fmt.Println("page size: ", os.Getpagesize())

	// Calculate nearest page size
	// QUESTION: I don't know how it works
	bufferSize = ((bufferSize + os.Getpagesize() - 1) / os.Getpagesize()) * os.Getpagesize()
	b.buflen = bufferSize
}

func (b *BipBuffer) Clear() {
	b.ixa = 0
	b.sza = 0
	b.ixb = 0
	b.szb = 0
	b.ixResrv = 0
	b.szResrv = 0
}

func (b *BipBuffer) FreeBuffer() {
	if b.buffer == nil {
		return
	}

	b.ixa, b.sza, b.ixb, b.szb, b.buflen = 0, 0, 0, 0, 0
	b.buffer = nil
}

// Reserve Reserves space in the buffer for a memory write operation
// Parameters:
//    int size                amount of space to reserve
//   OUT int& reserved        size of space actually reserved
//
// Returns:
//   BYTE*                    pointer to the reserved block
//
// Notes:
//   Will return NULL for the pointer if no space can be allocated.
//   Can return any value from 1 to size in reserved.
//   Will return NULL if a previous reservation has not been committed.
func (b *BipBuffer) Reserve(size int, reserved *int) int {
	if b.szb > 0 {
		var freespace = GetBFreeSpace()

		if size < freespace {
			freespace = size
		}

		if freespace == 0 {
			return 0
		}

		b.szResrv := freespace
		reserved = freespace
		ixResrv := b.ixb + b.szb
		return b.buflen + ixResrv
	} else {
		var freespace = GetBFreeSpace()
		if freespace >= b.ixa {
			if freespace == 0 {
				return 0
			}
			if size < freespace {
				freespace = size
			}

			b.szResrv = freespace
			reserved = &freespace
			b.ixResrv = b.ixa + b.sza
			return b.buflen + b.ixResrv

		} else {
			if b.ixa == 0 {
				return 0
			}

			if b.ixa < size {
				size = b.ixa
			}

			b.szResrv = size
			reserved = b.size
			b.ixResrv = 0
			return b.buflen
		}
	}
}

func GetBFreeSpace() int {
	return 0
}

func GetSpaceAfterA() int {
	return 0
}

func (b *BipBuffer) Commit(size int) {
	if size == 0 {
		b.szResrv = 0
		b.ixResrv = 0
		return
	}

	// if we are trying to commit more space than asked, then clip the size
	if size > b.szResrv {
		size = b.szResrv
	}

	// if we have no block currently, then we create on in A
	if b.sza == 0 && b.szb == 0 {
		b.ixa = b.ixResrv
		b.sza = size
		b.ixResrv = 0
		b.szResrv = 0
		return
	}

	if b.ixResrv == b.sza+b.ixa {
		b.sza += size
	} else {
		b.szb += size
	}

	b.ixResrv = 0
	b.szResrv = 0
}

func (b *BipBuffer) GetContiguousBlock(size *int) *int {
	if b.sza == 0 {
		size = new(int)
		return nil
	}

	size = &b.sza
	a := b.buflen + b.ixa
	return &a
}

func (b *BipBuffer) DecommitBlock(size int) {
	if size >= b.sza {
		b.ixa = b.ixb
		b.szResrv = b.szb
		b.ixb = 0
		b.szb = 0
	} else {
		b.sza -= size
		b.ixa -= size
	}
}

func (b *BipBuffer) GetCommittedSize() int {
	return b.sza + b.szb
}
func (b *BipBuffer) GetBufferSize() int {
	return b.buflen
}

func (b *BipBuffer) IsInitialized() bool {
	return b.buffer != nil
}

func (b *BipBuffer) GetSpaceAfterA() int {
	return b.buflen - b.ixa - b.sza
}

func (b *BipBuffer) GetBFreeSpace() int {
	return b.ixa - b.ixb - b.szb
}

func main() {

}