package main

import "fmt"

func main() {
  var a int32 = 10
  var b int32 = 10
  //var c rune = "d"
  // if the rune is given in double quotes it results in an error
  // cannot use "d" (type string) as type rune in assignment
  var c rune = '='
  fmt.Println("address of a " , &a)
  fmt.Println("address of b", &b)
  fmt.Println(" c " , c)

}
