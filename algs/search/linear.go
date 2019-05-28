// This file contains linear search algorithms.
// Author: Evan Douglass

package search

// Linear searches a slice of ints for the given value.
// It returns the index of the first occurrence of value if found, or -1 if not found.
func Linear(arr []int, value int) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}

// LinearSentinel searches a slice of ints for the given value using a
// sentinel value strategy.
// It returns the index of the first occurrence of value if found, or -1 if not found.
func LinearSentinel(arr []int, value int) int {
	length := len(arr)
	if length == 0 {
		return -1
	}

	// extract and replace last element
	last := arr[length-1]
	arr[length-1] = value
	// search for value
	i := 0
	for arr[i] != value {
		i++
	}
	// restore last element
	arr[length-1] = last

	if (i < length-1) || (arr[length-1] == value) {
		// found
		return i
	}
	// not found
	return -1
}

// LinearRecursive searches a slice of ints for the given value using a
// recursive strategy.
// It returns the index of the first occurrence of value if found, or -1 if not found.
func LinearRecursive(arr []int, i int, val int) int {
	if i >= len(arr) {
		return -1
	}
	if arr[i] == val {
		return i
	}
	return LinearRecursive(arr, i+1, val)
}
