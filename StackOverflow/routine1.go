package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	var x int
	threads := runtime.GOMAXPROCS(0) - 1
	log.Print(threads)
	for i := 0; i < threads; i++ {
		go func() {
			for {
				x++
				log.Print(x)
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("x =", x)
}
