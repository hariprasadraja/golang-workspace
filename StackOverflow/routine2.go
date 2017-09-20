package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	var x int
	threads := runtime.GOMAXPROCS(0)
	log.Print(threads)
	for i := 0; i < threads; i++ {
		go func() {
			for {
				x++
				time.Sleep(0)
				log.Print(x)
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("x =", x)
}
