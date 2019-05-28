// Package sort contains various sorting algorithms.
// These are academic exercises, not meant for use in production.
// As such, only integer slice operations are currently supported.
// Author: Evan Douglass
package sort

// SelectionSort sorts the given integer slice in-place.
func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		smallestI := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[smallestI] {
				smallestI = j
			}
		}
		arr[i], arr[smallestI] = arr[smallestI], arr[i]
	}
}
