package services

import (
    "errors"
    "fmt"

)

type Args struct {
    A int
    B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
    *reply = args.A * args.B
    fmt.Println("mmmmmmmmm")
    // mm := message{
    //     method: "broadcast",
    //     like:   3,
    // }
    //H.broadcast <- "hhh"
    return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
