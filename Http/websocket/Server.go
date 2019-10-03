package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)
		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		err := websocket.Message.Send(ws, msg)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	if err := http.ListenAndServe(":1234", websocket.Handler(Echo)); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
