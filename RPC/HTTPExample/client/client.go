package main

import (
	"golang-workspace/RPC/HTTPExample/service"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatalf("Error in dialing. %s", err)
	}
	args := &service.Args{
		A: 2,
		B: 3,
	}
	var result service.Result
	err = client.Call("Arith.Multiply", args, &result)
	if err != nil {
		log.Fatalf("error in Arith", err)
	}
	log.Printf("%d*%d=%d\n", args.A, args.B, result)
}
