package main

import (
	"bufio"
	"strings"
	"strconv"
	"fmt"
	"log"
)

func main() {

	// An artificial input source.
	const input = "1234 5678 1234567 9012345 67890"
	scanner := bufio.NewScanner(strings.NewReader(input))

     	// Create a custom split function by wrapping the existing ScanWords function.
	    split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		log.Println("token",token)
		log.Println("advance",advance)

		if err == nil && token != nil {
			n, _ := strconv.ParseInt(string(token), 10, 32)

		}
		return
	}
	// Set the split function for the scanning operation.
	scanner.Split(split)
	// Validate the input
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}
