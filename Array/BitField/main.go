package main

import "fmt"

/*
References:
1. https://en.wikipedia.org/wiki/Bit_field

*/

/* !!! Important !!!
Note: that signed bit-fields might not even be able to contain -1, this is true in computers that use 1-complement for representing signed int's. A signed bit-field that can contain 1 is a bug in the compiler.


I quote from WritingSolidCode, by SteveMaguire: "The bit field does have a non-zero state -- you just don't know what it is. The value can be either -1 or 1 [...]. You can safely use both states of the bit field if you restrict all your comparisons to 0."
--KatyMulvey
*/

/* Each of these preprocessor directives defines a single bit,
   corresponding to one button on the controller.  Button order
   matches that of the Nintendo Entertainment System. */
const (
	KEY_RIGHT  = 1 << iota /* 00000001   (1 << 0) */ //
	KEY_LEFT               /* 00000010    (1 << 0)  */
	KEY_DOWN               /* 00000100   (1 << 0) */
	KEY_UP                 /* 00001000  (1 << 0)  */
	KEY_START              /* 00010000 (1 << 0)  */
	KEY_SELECT             /* 00100000 (1 << 0)  */
	KEY_B                  /* 01000000 (1 << 0)  */
	KEY_A                  /* 10000000 (1 << 0)  */
)

var gameControllerStatus uint = 0

/* Sets the gameControllerStatus using OR */
func KeyPressed(key uint) {
	gameControllerStatus |= key
}

/* Turns the key in gameControllerStatus off using AND and ~ (binary NOT)*/
func KeyReleased(key uint) { gameControllerStatus &= ^key }

/* Tests whether a bit is set using AND */
func IsPressed(key uint) uint {
	return gameControllerStatus & key
}

func CheckNthBit(key uint, n uint) uint {
	return (key >> n) & 1
	// return key & (1 << n) != 0
}

func main() {
	fmt.Println(KEY_RIGHT)
	fmt.Println(KEY_LEFT)
	fmt.Println(KEY_DOWN)
}
