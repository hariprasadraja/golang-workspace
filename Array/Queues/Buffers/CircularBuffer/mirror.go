package main

import "errors"

func allocate_mirrored(howmuch uint, howmany uint) error {
	if howmuch <= 0 || howmuch != (howmuch*GetPageSize()*GetPageSize()) {
		return errors.New("Invalid")
	}

	// mem variable will hold the allocated pointer
	var mem interface{}

}

func CheckError() {

}
