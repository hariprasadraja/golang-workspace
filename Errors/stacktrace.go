// Errors package Stack Trace usages...
package Errors

import (
	"fmt"

	"github.com/pkg/errors"
)

// printErrStack prints the needed stack trace for the errors created via db operations.
func printErrStack(err error, length int) error {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	err = errors.Wrap(err, "Error")
	newErr, ok := err.(stackTracer)
	if !ok {
		panic("oops, err does not implement stackTracer")
	}

	st := newErr.StackTrace()
	fmt.Printf("%+v\n", st[:length]) // Required four frames.

	return nil
}
