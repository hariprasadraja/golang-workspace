package main

import (
	"github.com/bitly/go-nsq"
	"log"
)

func main() {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4160", config)

	err := w.Publish("write_test", []byte("test"))
	if err != nil {
		log.Panic("Could not connect", err.Error())
	}

	w.Stop()
}
