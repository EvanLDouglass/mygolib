package search

import "testing"

func TestLinear(t *testing.T) {
	cases := []struct {
		inArr    []int
		inVal    int
		expected int
	}{
		// tests
		{
			[]int{}, 3, -1,
		},
		{
			[]int{1}, 1, 0,
		},
		{
			[]int{1}, 2, -1,
		},
		{
			[]int{1, 2, 3, 4, 5}, 3, 2,
		},
		{
			[]int{4, 3, 7, 1, 5, 0, 23, 90, 1000, 2, 5}, 5, 4,
		},
	}
	for _, c := range cases {
		// test Linear
		got := Linear(c.inArr, c.inVal)
		if got != c.expected {
			t.Errorf("Linear(%v, %v) == %v, expected %v", c.inArr, c.inVal, got, c.expected)
		}

		// test LinearSentinel
		got = LinearSentinel(c.inArr, c.inVal)
		if got != c.expected {
			t.Errorf("LinearSentinel(%v, %v) == %v, expected %v", c.inArr, c.inVal, got, c.expected)
		}

		// test LinearRecursive
		got = LinearRecursive(c.inArr, 0, c.inVal)
		if got != c.expected {
			t.Errorf("LinearRecursive(%v, 0, %v) == %v, expected %v", c.inArr, c.inVal, got, c.expected)
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	arr := make([]int, 10000)
	val := 1 // will not be found
	for i := 0; i < b.N; i++ {
		Linear(arr, val)
	}
}

func BenchmarkSentinelSearch(b *testing.B) {
	arr := make([]int, 10000)
	val := 1 // will not be found
	for i := 0; i < b.N; i++ {
		LinearSentinel(arr, val)
	}
}

func BenchmarkRecursiveSearch(b *testing.B) {
	arr := make([]int, 10000)
	val := 1 // will not be found
	for i := 0; i < b.N; i++ {
		LinearRecursive(arr, 0, val)
	}
}

func BenchmarkLinearSearchLarge(b *testing.B) {
	arr := make([]int, 1000000)
	val := 1 // will not be found
	for i := 0; i < b.N; i++ {
		Linear(arr, val)
	}
}

func BenchmarkSentinelSearchLarge(b *testing.B) {
	arr := make([]int, 1000000)
	val := 1 // will not be found
	for i := 0; i < b.N; i++ {
		LinearSentinel(arr, val)
	}
}

func BenchmarkLinearRecursiveSearchLarge(b *testing.B) {
	arr := make([]int, 1000000)
	val := 1 // will not be found
	for i := 0; i < b.N; i++ {
		LinearRecursive(arr, 0, val)
	}
}
