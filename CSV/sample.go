package main

import (
	"encoding/csv"
	"log"
	"os"
)

var data = []string{"Reason", "Request", "Error"}

func main() {
	file, err := os.Create("result.csv")
	defer file.Close()
	if err != nil {
		log.Fatal("Failed to create csv file, Error: ", err.Error())
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(data)
	if err != nil {
		log.Fatal("Failed to write into the file, Error: ", err.Error())
	}
}
