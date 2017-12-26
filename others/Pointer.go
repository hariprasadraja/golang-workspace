package main

import "github.com/emicklei/go-restful/log"

func main() {

	var newvar *int
	var data int  = 10
	newvar = &data
	log.Print(*newvar)
}
