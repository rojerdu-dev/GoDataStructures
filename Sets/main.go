package main

import (
	"errors"
	"fmt"
)

type Set struct {
	Items map[string]struct{}
}

// Create new set
func NewSet() *Set {
	var set Set
	set.Items = make(map[string]struct{})
	return &set
}

// Add item to set
func (s *Set) Add(item string) {
	s.Items[item] = struct{}{}
}

// Remove item from set
func (s *Set) Remove(item string) error {
	if _, exists := s.Items[item]; !exists {
		return errors.New("item not found in set")
	}
	delete(s.Items, item)
	return nil
}

// Contains item in set
func (s *Set) Contains(item string) bool {
	_, exists := s.Items[item]
	return exists
}

func (s *Set) List() {
	for k, _ := range s.Items {
		fmt.Println(k)
	}
}

func main() {

	programming := NewSet()

	programming.Add("C")
	programming.Add("SQL")
	programming.Add("C++")
	programming.Add("Java")
	programming.Add("Python")
	programming.Add("Ruby")
	programming.Add("JavaScript")
	programming.Add("Go")

	programming.List()

	exists := programming.Contains("Haskell")
	fmt.Printf("Haskell exists in the set: %t\n", exists)
}
