package Errors

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
)

// printErrStack prints the needed stack trace for the errors created via db operations.
func printErrStack(err error) error {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	err = errors.Wrap(err, "Error")
	log.Print(err)
	newErr, ok := err.(stackTracer)
	if !ok {
		panic("oops, err does not implement stackTracer")
	}

	st := newErr.StackTrace()
	fmt.Printf("%+v\n", st[1:3]) // Required four frames.

	return nil
}
