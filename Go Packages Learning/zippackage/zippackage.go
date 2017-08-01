package main

import (
	"archive/zip"
	"log"
	"fmt"
	"io"
	"os"
)

func main() {
	// Open a zip archive for reading.
	r, err := zip.OpenReader("/home/parvathavarthinik/learn.zip")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(r.File[0].Name)
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}

}