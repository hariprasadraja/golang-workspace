package main

import (
	"github.com/pkg/errors"
)

type DBError struct {
	code    string
	message string
}

type ValidationErr struct {
	code    string
	message string
	Errors  *ValidationErr
}

//type Error struct {
//	code string
//	message string
//	Error *ValidationErr
//}

// Returns the Type of Error either
// DB or Validation
func ErrHandler(err interface{}) error {
	switch err {
	case err.(*DBError):
		{
			DBErr := err.(*DBError)
			error := errors.Errorf("%s", DBErr.message)
			return error
		}

	case err.(*ValidationErr):
		{
			ValidErr := err.(*ValidationErr)
			error := errors.Errorf("%s", ValidErr.message)
			return error
		}
	}
	return nil
}

func main() {
	err := &DBError{}
	err.code = "10"
	ErrHandler(err)

	valErr := &ValidationErr{}
	valErr.code = "20"
	ErrHandler(err)
}
