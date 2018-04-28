package main

import "fmt"

// func closuretest(a string) func(s string)  {
//   return func(s string) {
//     fmt.Println("inside the function", a,s)
//   }
// }
type name func(int, int) int

func Calculator(operator name, val1 int, val2 int) int {
	return operator(val1, val2)
}

func main() {
	//closuretest("Hi i am hari")
	//  name := closuretest("hi")
	// name("hari")
	//   func(a string) {
	//   fmt.Println(a)
	// }("hi")

	// Passing a function as a parameter to another function
	// add:= func(int a, int b) int {
	//   return a+b
	// }
	// sub:= func(int a,int b) int {
	//   return a-b
	// }
	// fmt.Println(sub(2,3))
	//  operation:= Calculator(add,40,50)
	//  fmt.Println(operation)
	//  operation = Calculator(sub,10,5)
	//   fmt.Println(operation)

	//Annoymous function retruns  the parameter for another funciton.
	//panic("Something went wrong")
	diff := Calculator(func(a int, b int) int {
		return a + b
	}, 10, 5)
	fmt.Println(diff)

}
