package main

import "github.com/emicklei/go-restful/log"

func main() {
	var pipe chan map[string]string
	pipe = make(chan map[string]string, 2)
	go connect("myhost", "100", pipe)
	out := <-pipe

	log.Print(out)

}

func connect(host string, url string, pipe chan<- map[string]string) {
	log.Print("Trying " + url)
	var lpipe map[string]string
	lpipe = make(map[string]string)
	lpipe["resp"], lpipe["err"] = "aaa", "bbb"
	pipe <- lpipe
}
