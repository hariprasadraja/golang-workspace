package main

import (
	"fmt"
	"github.com/gopherjs/gopherjs/compiler/natives/src/sync"
	"log"
	"net/http"
	"os"
	//"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	url := "http://localhost:9000/"
	logFile, err := os.OpenFile("log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		fmt.Printf("Error in openning log file: %s", err.Error())
		os.Exit(1)
	}

	go call(url, "First:", logFile)
	go call(url, "Second:", logFile)
	go call(url, "Third:", logFile)
	go call(url, "Fourth:", logFile)
	go call(url, "Fifth:", logFile)
	wg.Wait()

}

func call(url string, routine string, logfile *os.File) {
	i := 0
	for {
		_, err := http.Get(url)
		if err != nil {
			log.Println("error", err)

		}
		//time := time.Now().Nanosecond()
		//req,err := http.NewRequest("GET",url+fmt.Sprint(time),nil)
		//if err != nil {
		//	log.Println("error 0",err)
		//}
		//
		//c := &http.Client{}
		//
		//_,err = c.Do(req)
		//if err != nil {
		//	log.Println("error 1",err)
		//}
		i++
		log.Println(routine, i)
		log.SetOutput(logfile)
	}
}
