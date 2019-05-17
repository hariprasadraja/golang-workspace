package datastructures

import (
	"fmt"
	"strings"
)

// Set (data structure) stores unique elements in it.
type Set struct {
	Items map[interface{}]bool
}

// Adds the new values in to the Set
func (s *Set) Add(values ...interface{}) {
	if s.Items == nil {
		s.Items = make(map[interface{}]bool)
	}

	for _, value := range values {
		s.Items[value] = true
	}
}

// Delete removes value in the Set
func (s *Set) Delete(values ...interface{}) {
	for _, value := range values {
		delete(s.Items, value)
	}
}

// Has returns true if the Set contains any snigle Item present
func (s *Set) HasANY(values ...interface{}) ([]interface{}, bool) {
	matched := []interface{}{}
	anyMatch := false

	for _, value := range values {
		if _, ok := s.Items[value]; ok {
			matched = append(matched, value)
			anyMatch = true
		}
	}

	return matched, anyMatch
}

// HasALL returns true, if it all the values are present in the array, if some values are not in the set
// then it will return those values in not matched.
func (s *Set) HasALL(values ...interface{}) ([]interface{}, bool) {
	notMatched := []interface{}{}
	allMatched := true

	for _, value := range values {
		if _, ok := s.Items[value]; !ok {
			allMatched = false
			notMatched = append(notMatched, value)
		}
	}

	return notMatched, allMatched
}

// Next returns the Next value in the set to read
// Example
//      for key,val := range set.Next()
func (s *Set) Next() <-chan interface{} {
	newChan := make(chan interface{})
	go func() {
		for key, _ := range s.Items {
			newChan <- key
		}

		close(newChan)
	}()

	return newChan
}

// Clear removes all elements from the Set
func (s *Set) Clear() {
	s.Items = make(map[interface{}]bool)
}

// String prints the Set in human readable format
func (s *Set) String() string {
	var printStr string
	for key, _ := range s.Items {
		printStr += fmt.Sprintf("%v,", key)
	}

	return "{" + strings.TrimSuffix(printStr, ",") + "}"
}
