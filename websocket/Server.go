package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"log"
	"net/http"
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

		 err := websocket.Message.Send(ws,msg)
		 if(err != nil) {
			 panic(err)
		 }
	}
}



func main() {

	if err := http.ListenAndServe(":1234", websocket.Handler(Echo)); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}