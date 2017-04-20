package main

import (
	"os"
	"log"
	"archive/tar"
	"io"
)

func main() {
 // tar package is mainly focused on providing customised feel of creating and writing tar files
 // the Header struct in the tar file is the major part of the tar file which gives us the customisation over the tar file.
	// it also provides read and write operations wrapped around the header of the tar file

    data := make([]byte,100)
	file,err :=os.Open("/home/parvathavarthinik/learn.tar")
	if err != nil{
		panic(err)
	}
	log.Println(file)
	file.Read(data)
	s :=string(data)
	log.Println(s)


	Reader :=tar.NewReader(file)
	Reader.Read(data)
	log.Println(string(data))

	hdr := &tar.Header{
		Name:"updatedfile",
	}

	 pen :=tar.NewWriter(file)
	 pen.WriteHeader(hdr)
	 pen.Write(data)
	 defer pen.Flush()
}
