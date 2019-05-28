package search

import "testing"

var none = []int{}
var one = []int{1}
var two = []int{1, 3}
var three = []int{2, 4, 6}
var many = []int{1, 1, 2, 4, 5, 5, 5, 6, 10, 22, 22, 80, 90, 90, 90, 90, 90}

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		inArr    []int
		inVal    int
		expected int
	}{
		// tests
		{none, 0, -1},
		{one, 1, 0},
		{one, 9, -1},
		{two, 1, 0},
		{two, 3, 1},
		{two, 5, -1},
		{three, 2, 0},
		{three, 4, 1},
		{three, 6, 2},
		{three, 8, -1},
		{many, 10, 8},
		{many, 1, 1},
		{many, 90, 12},
		{many, 2000, -1},
	}

	for _, c := range cases {
		got := Binary(c.inArr, c.inVal)
		if got != c.expected {
			t.Errorf("Binary(%v, %v) == %v, expected %v",
				c.inArr, c.inVal, got, c.expected)
		}
	}
}

func TestBinaryRecursiveSearch(t *testing.T) {
	cases := []struct {
		inArr    []int
		inVal    int
		expected int
	}{
		// tests
		{none, 0, -1},
		{one, 1, 0},
		{one, 9, -1},
		{two, 1, 0},
		{two, 3, 1},
		{two, 5, -1},
		{three, 2, 0},
		{three, 4, 1},
		{three, 6, 2},
		{three, 8, -1},
		{many, 10, 8},
		{many, 1, 1},
		{many, 90, 12},
		{many, 2000, -1},
	}

	for _, c := range cases {
		got := BinaryRecursive(c.inArr, c.inVal)
		if got != c.expected {
			t.Errorf("BinaryRecursive(%v, %v) == %v, expected %v",
				c.inArr, c.inVal, got, c.expected)
		}
	}
}

func BenchmarkBinarySearchSmall(b *testing.B) {
	arr := make([]int, 1000) // filled with zeros
	for i := 0; i < b.N; i++ {
		Binary(arr, 1)
	}
}

func BenchmarkBinaryRecursiveSearchSmall(b *testing.B) {
	arr := make([]int, 1000)
	for i := 0; i < b.N; i++ {
		BinaryRecursive(arr, 1)
	}
}

func BenchmarkBinarySearchLarge(b *testing.B) {
	arr := make([]int, 1000000) // filled with zeros
	for i := 0; i < b.N; i++ {
		Binary(arr, 1)
	}
}

func BenchmarkBinaryRecursiveSearchLarge(b *testing.B) {
	arr := make([]int, 1000000) // filled with zeros
	for i := 0; i < b.N; i++ {
		BinaryRecursive(arr, 1)
	}
}
