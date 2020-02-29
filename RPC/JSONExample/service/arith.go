// package service will have all the RPC apis required for the RPC client and the server communication.

package service

import (
	"log"
	"net/http"
)

//Holds arguments to be passed to service Arith in RPC call
type Args struct {
	A, B int
}

type Arith int

type Result int

func (t *Arith) Multiply(r *http.Request, args *Args, result *Result) error {
	log.Printf("Multiplying %d with %d\n", args.A, args.B)
	*result = Result(args.A * args.B)
	return nil
}
