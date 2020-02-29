package service

import (
	"errors"
	"log"
)

type Arith int

type MultiplyReq struct {
	A, B int
}
type MultiplyRes int

func (t *Arith) Multiply(args MultiplyReq, reply *MultiplyRes) error {
	log.Print("Multiplying %+v", args)
	*reply = MultiplyRes(args.A * args.B)
	return nil
}

type DivideReq struct {
	A, B int
}

type DivideRes struct {
	Quo, Rem int
}

func (t *Arith) Divide(args DivideReq, quo *DivideRes) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
