package main

import (
	"golang-workspace/RPC/TCPExample/service"
	"log"
	"net/rpc"
)

func main() {

	// we can also use net.Dail instead
    client,err := rpc.Dial("tcp",":1234")
	handleErr(err)

	args := service.MultiplyReq{
		A:5,
		B:6,
	}

	var result service.MultiplyRes = 0
	err = client.Call("Arith.Multiply",args,&result)
	handleErr(err)

	print(result)
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}