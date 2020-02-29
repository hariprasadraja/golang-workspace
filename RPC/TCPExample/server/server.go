package main

import (
	"golang-workspace/RPC/TCPExample/service"
	"log"
	"net"
	"net/rpc"
)

func main() {
	arithService := new(service.Arith)
	err := rpc.RegisterName("Arith", arithService)
	if err != nil {
		log.Fatal(err)
	}

	// here you can use either http or tcp or any other connection you want
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	rpc.Accept(listener)
	defer listener.Close()

	// you can also try writing like this explicitly calling the Accept method of the net.listener
	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	rpc.ServeConn(conn)
	// }

}
