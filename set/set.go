// Author: Evan Douglass
// Created On: April 30 2019

// Package set implements the functionality of a set.
// The package is inspired from an article in the go blog
// about maps, which is the underlying data structure for a Set.
// A Set has constant lookup time, but linear time for most
// operations involving sets of more than one element.
// As implemented here, a Set can contain any type that can be a map key.
// It can also hold different types in the same set.
package set

import (
	"fmt"
	"reflect"
	"strings"
)

// Set is the primary data structure of this package.
// A set can hold only one of any unique value.
// Values can be of any type that can be a map key.
type Set struct {
	set map[interface{}]bool
}

// NewSet creates a set and returns a pointer to it.
// Allows setting a variadic number of initial elements.
// Go will not allow unpacking of slices unless first made into
// type []interface{}.
func NewSet(initVals ...interface{}) *Set {
	set := Set{
		set: make(map[interface{}]bool),
	}

	for _, v := range initVals {
		set.AddVal(v)
	}

	return &set
}

// HasVal determines if a set contains the given value
func (s *Set) HasVal(val interface{}) bool {
	return s.set[val]
}

// IsEmpty determines if a set has no elements
func (s *Set) IsEmpty() bool {
	return s.Len() == 0
}

// AddVal adds an element to a set.
func (s *Set) AddVal(val interface{}) {
	s.set[val] = true
}

// RemoveVal removes an element from a set. It does nothing if the set does
// not contain the value.
func (s *Set) RemoveVal(val interface{}) {
	delete(s.set, val)
}

// Union returns the union of the given sets as a new Set.
func (s *Set) Union(s2 *Set) *Set {
	new := NewSet() // Pointer to new set

	// Add all vals in s to new set
	for k := range s.set {
		new.AddVal(k)
	}

	// Add vals from s2 not already in new set
	for k := range s2.set {
		// Will reassign existing vals instead of adding duplicates
		new.AddVal(k)
	}

	return new
}

// Inter returns the intersection of the given sets as a new Set.
func (s *Set) Inter(s2 *Set) *Set {
	new := NewSet()
	for k := range s.set {
		// Add all values in s that are also in s2
		if s2.HasVal(k) {
			new.AddVal(k)
		}
	}
	return new
}

// Diff returns the difference of the two sets (s - s2) as a new Set.
func (s *Set) Diff(s2 *Set) *Set {
	new := NewSet()
	for k := range s.set {
		// Add all values in s that are not in s2
		if !s2.HasVal(k) {
			new.AddVal(k)
		}
	}
	return new
}

// Len returns the length of a set.
func (s *Set) Len() int {
	return len(s.set)
}

// Equals determines if two sets contain the same elements.
func (s *Set) Equals(s2 *Set) bool {
	return reflect.DeepEqual(s.set, s2.set)
}

// String represents a Set as a string.
// If the set contains an int and float that are the same number (i.e. 1 and 1.0),
// or a string that resembles another element (i.e. "1" and 1 or 1.0), they
// will look the same in the printed string.
func (s *Set) String() string {
	var b strings.Builder

	// Add identifier
	b.WriteString("Set( ")

	// Add elements if any
	for k := range s.set {
		fmt.Fprintf(&b, "%v ", k)
	}

	// End string
	b.WriteString(")")
	return b.String()
}
