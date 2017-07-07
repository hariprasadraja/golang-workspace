package main

import (
	"fmt"
	"strings"
)

func main() {
	var data []string
	x := struct {
		Foo string
		Bar int
	}{"foo", 2}
	s := fmt.Sprintf("%+v", x)
	fmt.Println(s)
	s = strings.Replace(s, "{", "", -1)
	s = strings.Replace(s, "}", "", -1)
	data = strings.Split(s, ":")
	fmt.Println(data)
	for _, val := range data {
		fmt.Println(val)
	}

}
