package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

/*
Set (data structure) stores unique elements in it.
by default, Set is not safe for Concurrency.
Example
	safeSet := new(Set).Safe()
	notSafeSet := new(Set)
*/
type Set struct {
	items          map[interface{}]bool
	rwLock         sync.RWMutex
	concurrentSafe bool // set true, if the Set should be concurrent safe
}

// Adds the new values in to the Set
func (s *Set) Add(values ...interface{}) {
	s.lock()
	defer s.unlock()

	if s.items == nil {
		s.items = make(map[interface{}]bool)
	}

	for _, value := range values {
		s.items[value] = true
	}
}

// Delete removes value in the Set
func (s *Set) Delete(values ...interface{}) {
	s.lock()
	defer s.unlock()

	for _, value := range values {
		delete(s.items, value)
	}
}

// HasMatch compares the values with the set returns all the Matched and Unmatched values in the set
func (s *Set) HasMatch(values ...interface{}) (matched []interface{}, Unmatched []interface{}) {
	s.rlock()
	defer s.rUnlock()

	matched = make([]interface{}, 0)
	Unmatched = make([]interface{}, 0)
	for _, value := range values {
		if _, ok := s.items[value]; ok {
			matched = append(matched, value)
		} else {
			Unmatched = append(Unmatched, value)
		}
	}

	return
}

// HasOne returns true, if the value is present in the set
func (s *Set) HasOne(value interface{}) bool {
	s.rlock()
	defer s.rUnlock()

	_, ok := s.items[value]
	return ok
}

// HasAll returns true, if all the values are present in the set
func (s *Set) HasAll(values ...interface{}) bool {
	s.rlock()
	defer s.rUnlock()

	for _, value := range values {
		if _, ok := s.items[value]; !ok {
			return false
		}
	}

	return true
}

// SubSet returns true, if the set is the subset of set 's'
func (s *Set) SubSet(set *Set) bool {
	s.rlock()
	set.rlock()
	defer s.rUnlock()
	defer set.rUnlock()

	for value := range set.Next() {
		if _, ok := s.items[value]; !ok {
			return false
		}
	}

	return true
}

// Next returns the Next value in the set to read
// Example
//      for key,val := range set.Next()
func (s *Set) Next() <-chan interface{} {
	s.rlock()
	defer s.rUnlock()

	newChan := make(chan interface{})
	go func() {
		for key, _ := range s.items {
			newChan <- key
		}

		close(newChan)
	}()

	return newChan
}

// Clear removes all elements from the Set
func (s *Set) Clear() {
	s.lock()
	defer s.unlock()
	s.items = make(map[interface{}]bool)
}

// String prints the Set in human readable format
func (s *Set) String() string {
	s.rlock()
	defer s.rUnlock()

	var printStr string
	for key, _ := range s.items {
		printStr += fmt.Sprintf("%v,", key)
	}

	return "{" + strings.TrimSuffix(printStr, ",") + "}"
}

// Size returns the size of the set
func (s *Set) Size() int {
	s.rlock()
	defer s.rUnlock()
	return len(s.items)
}

// Items returns all items in the set
func (s *Set) Items() []interface{} {
	s.rlock()
	defer s.rUnlock()

	items := make([]interface{}, 0)
	for i := range s.items {
		items = append(items, i)
	}

	return items
}

// Union retruns a new set with elements from both the Sets
// new set is not concurrent safe
func (s *Set) Union(set *Set) (newSet *Set) {
	s.rlock()
	set.rlock()
	defer s.rUnlock()
	defer set.rUnlock()

	newSet = new(Set)
	newSet.Add(s.Items()...)
	newSet.Add(set.Items()...)
	return newSet
}

// Intersection returns a new set with elements that exist in both sets
// new set is not concurrent safe
func (s *Set) Intersection(set *Set) (newSet *Set) {
	s.rlock()
	set.rlock()
	defer s.rUnlock()
	defer set.rUnlock()

	matched, _ := s.HasMatch(set.Items()...)
	newSet = new(Set)
	newSet.Add(matched...)
	return newSet
}

// Difference returns a new set with all the elements that
// new set is not concurrent safe
func (s *Set) Diffrence(set *Set) (newSet *Set) {
	s.rlock()
	set.rlock()
	defer s.rUnlock()
	defer set.rUnlock()

	newSet = new(Set)
	for i := range s.Next() {
		if !set.HasOne(i) {
			newSet.Add(i)
		}
	}

	return newSet
}

func (s *Set) lock() {
	if s.concurrentSafe {
		s.rwLock.Lock()
	}
}

func (s *Set) unlock() {
	if s.concurrentSafe {
		s.rwLock.Unlock()
	}
}

func (s *Set) rlock() {
	if s.concurrentSafe {
		s.rwLock.RLock()
	}
}

func (s *Set) rUnlock() {
	if s.concurrentSafe {
		s.rwLock.RUnlock()
	}
}

// Safe returns the give Set as ConcurrentSafe
func (s *Set) Safe() *Set {
	s.concurrentSafe = true
	return s
}

// UnSafe returns the given Set as non Concurrent Safe
func (s *Set) UnSafe() *Set {
	s.concurrentSafe = false
	return s
}

func main() {
	mySet := new(Set).Safe()
	mySet.Add(1, 2, 3, "hi", 4.0, 6.0, true)

	secondSet := new(Set)
	secondSet.Add(100)

	// TEMP: Snippet for debugging. remove it before commit
	fmt.Printf("\n set :: %v \n\n", mySet)

	// TEMP: Snippet for debugging. remove it before commit
	fmt.Printf("\n secondSet :: %+v \n\n", secondSet)

	go mySet.Add(1, 2, 3, "hi", 4.0, 6.0, true)
	go mySet.Add(4, 5, 6, 7)
	go mySet.Add("A", "B", "C", "D")
	go mySet.Delete("a", "b", "false")
	go mySet.Delete("a", "b", "false")
	go mySet.Add("a", "b", "c", "hello", false)
	go mySet.Add("a", "b", "c", "hello", false)
	go mySet.Add("a", "b", "c", "hello", false)
	go mySet.Delete("a", "b", "false", false)
	go mySet.Add("a", "b", "c", "hello", false)
	go mySet.Delete("a", "b", "false")

	time.Sleep(10 * time.Second)

	// TEMP: Snippet for debugging. remove it before commit
	fmt.Printf("\n set :: %v \n\n", mySet)

	// TEMP: Snippet for debugging. remove it before commit
	fmt.Printf("\n len(set) :: %v \n\n", mySet.Size())

	fmt.Printf("\n subset :: %v \n\n", mySet.SubSet(secondSet))

}