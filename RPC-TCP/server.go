package main

import "net/rpc"

func main() {
	arithService := new(Arith)
	rpc.RegisterName("arith", arithService)

	// here you can use either http or tcp or any other connection you want
	rpc.ServeConn()

}
