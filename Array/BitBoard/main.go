package main

import "fmt"

type BitBoard uint

func SetCell(bitboard BitBoard, row uint, col uint) BitBoard {
	newbit := BitBoard(1 << (row*8 + col))
	return bitboard | newbit
}

func GetCell(bitboard BitBoard, row uint, col uint) bool {
	mask := BitBoard(1 << (row*8 + col))
	return (bitboard & mask) != 0
}

// BitsCount is to count the number of bits turned on in the bit board
func BitsCount(bitboard BitBoard) uint {
	var count uint
	for bitboard != 0 {

		bitboard &= bitboard - 1
		count += 1
	}

	return count
}

func main() {
	var board BitBoard

	newBoard := SetCell(board, 3, 5)
	fmt.Println(newBoard)

	if GetCell(newBoard, 3, 5) {
		fmt.Println("Bit is set in the cell")
	} else {
		fmt.Println("Bit is not set in the cell")
	}

	fmt.Println("Old Board Bits Count :", board)
	fmt.Println("New Board Bits Count:", BitsCount(newBoard))
}
