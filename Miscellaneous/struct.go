// Struct in google go
package main

import "fmt"

type test struct {
	a int
	b float32
}

type test1 struct {
	a int
	b int
}

func (ref4 test1) mytest() int {
	return (ref4.a) * (ref4.b)
}
func main() {
	/* Case 1 : assaigning value to the struc using reference*/
	var ref1 test
	ref1.a = 10
	ref1.b = 20
	fmt.Println(ref1)
	ref1.a = 40
	ref1.b = 40
	/*Case 2: assaigning value to the struct at the time of declaring the reference*/
	ref2 := test1{a: 10, b: 20}
	fmt.Println(ref2)
	/* Case 3: Creating the pointer variable for the struct */
	ref3 := new(test)
	ref3.a = 20
	ref3.b = 10
	fmt.Println(*ref3)
	fmt.Println(ref2.mytest())
}
