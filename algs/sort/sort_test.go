package sort

import (
	"math/rand"
	"testing"
)

var (
	none  = make([]int, 0)
	one   = make([]int, 1)
	two   = make([]int, 2)
	three = make([]int, 3)
	many  = make([]int, 18)
)

func setUp() {
	none = []int{}
	one = []int{2}
	two = []int{3, 1}
	three = []int{3, 5, 1}
	many = []int{1, 1, 4, 6, 2, 4, 5, 8, 100, 454, 22, 1, 1, 45, 99, -15, 33, 44}
	// sorted: -15, 1, 1, 1, 1, 2, 4, 4, 5, 6, 8, 22, 33, 44, 45, 99, 100, 454
}

// Equal provides a test of equality for integer slices
func Equal(xi1 []int, xi2 []int) bool {
	if len(xi1) != len(xi2) {
		return false
	}

	for i := range xi1 {
		if xi1[i] != xi2[i] {
			return false
		}
	}
	return true
}

/* ===== Tests ===== */

func TestSelectionSort(t *testing.T) {
	setUp()
	cases := []struct {
		in       []int
		expected []int
	}{
		{none, []int{}},
		{one, []int{2}},
		{two, []int{1, 3}},
		{two, []int{1, 3}}, // can sort a sorted slice?
		{three, []int{1, 3, 5}},
		{many, []int{-15, 1, 1, 1, 1, 2, 4, 4, 5, 6, 8, 22, 33, 44, 45, 99, 100, 454}},
	}

	for _, c := range cases {
		SelectionSort(c.in)
		if !Equal(c.in, c.expected) { // in-place sorting
			t.Errorf("\n=== Selection Sort ===\nGot:     %v\nExpected:%v", c.in, c.expected)
		}
	}
}

/* ===== Benchmarks ===== */

func BenchmarkSetUp(b *testing.B) {
	var arr = make([]int, 50000)
	for i := 0; i < b.N; i++ {
		// populate arr with random numbers
		for i := range arr {
			arr[i] = rand.Intn(1000)
		}
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	var arr = make([]int, 50000)
	for i := 0; i < b.N; i++ {
		// populate arr
		for i := range arr {
			arr[i] = rand.Intn(1000)
		}
		// sort it
		SelectionSort(arr)
	}
}
