// Package binary contains binary search algorithms.
// These are academic exercises, not meant for use in production.
// Author: Evan Douglass
package binary

// NOT_FOUND represents the return value when a search is unsuccessful.
const NOT_FOUND int = -1

// Search conducts a binary search on a sorted array. If the array given is not
// sorted, bad things will happen!
// Returns the index of the first found occurrence of val, or -1 if not found.
func Search(arr []int, val int) int {
	p := 0
	r := len(arr) - 1

	for p <= r {
		mid := (p + r) / 2
		if arr[mid] == val {
			return mid
		} else if arr[mid] > val {
			r = mid - 1
		} else {
			p = mid + 1
		}
	}

	return NOT_FOUND
}

// RecursiveSearch conducts a recursive binary search on a sorted array.
// If the array given is not sorted, bad things will happen!
// Returns the index of the first found occurrence of val, or -1 if not found.
func RecursiveSearch(arr []int, val int) int {
	end := len(arr) - 1
	return recursiveSearchHelper(arr, 0, end, val)
}

// Helper for RecursiveSearch
func recursiveSearchHelper(arr []int, p int, r int, val int) int {
	if p > r {
		return NOT_FOUND
	}

	mid := (p + r) / 2
	if arr[mid] == val {
		return mid
	} else if arr[mid] < val {
		return recursiveSearchHelper(arr, mid+1, r, val)
	} else {
		return recursiveSearchHelper(arr, p, mid-1, val)
	}
}
