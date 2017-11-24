
package main

import "fmt"
func main() {
  var a [5]int
  a[0] = 0
  a[1] = 1
  a[2] = 2
  a[3] = 3
  a[4] = 4
  for sum:= range a {
    fmt.Println(a[sum])
    // checking Println is a vardic funtion

    fmt.Println(a[sum],a[sum],"hi")
  }
  var g, h, i int = 5, 10, 20
  fmt.Println(g)
e, b, c := 80, 80, 80
fmt.Println(e,b,c)
}
