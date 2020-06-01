package main

import (
	"fmt"
)

// References:
// https://en.wikipedia.org/wiki/Bit_array
// https://github.com/golang-collections/go-datastructures/blob/master/bitarray/block.go

// 	&    bitwise AND            integers
// 	|    bitwise OR             integers
// 	^    bitwise XOR            integers
// 	&^   bit clear (AND NOT)    integers

// BitArray represents a value as bits.
// []bool does not support bit wise operations in golang. So uint32 has been used to represent the bit.
type BitArray uint32

// NOT is used to invert all the bits in an array
func NOT(array BitArray) BitArray {
	return ^array
}

func AND(array1, array2 BitArray) BitArray {
	return array1 & array2
}

func InsertBit(bitArray BitArray, position uint64) BitArray {
	return bitArray | BitArray(1<<position)
}

func RemoveBit(bitArray BitArray, position uint64) BitArray {
	return bitArray & ^(1 << position)
}

func GetBit(bitArray BitArray, position uint64) bool {
	return bitArray&BitArray(1<<position) != 0
}

func Equals(bitArray BitArray, other BitArray) bool {
	return bitArray == other
}

func Intersects(bitArray BitArray, other BitArray) bool {
	return bitArray&other == other
}

func main() {
	var block BitArray

	block = InsertBit(block, 4)
	fmt.Println(block) // 0000 1000  -> 16

	fmt.Println(RemoveBit(block, 4))

	fmt.Println(AND(block, 7))
	fmt.Print(NOT(block)) // 4294967279 prints the full capacity of the array.

}
