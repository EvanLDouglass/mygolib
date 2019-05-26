// Tests for the set package
package set

import (
	"fmt"
	"testing"
)

func TestNewSet(t *testing.T) {
	set1 := NewSet()
	set2 := Set{make(map[interface{}]bool)}

	if !set1.IsEmpty() {
		t.Errorf("Expected an empty set, got a set with %v elements", set1.Len())
	}

	equal := set1.Equals(&set2)
	if !equal {
		t.Errorf("Expected empty sets to be equal, got set1=%v, set2=%v", *set1, set2)
	}

	set1 = NewSet(1, "Hello", 3.1415)
	if set1.Len() != 3 {
		t.Error("expected set1.Len()==3, got", set1.Len())
	}

	if !set1.HasVal(1) || !set1.HasVal("Hello") || !set1.HasVal(3.1415) {
		t.Error("test failed, got set", set1)
	}
}

func TestLen(t *testing.T) {
	set := NewSet()

	// Test zero length
	if !set.IsEmpty() {
		t.Errorf("Expected len=0, got len=%v\n%v", set.Len(), set.set)
	}

	// Add one
	set.AddVal("test1")
	if l := set.Len(); l != 1 {
		t.Error("Expected len=1, got len =", l)
	}

	// Add four more
	set.AddVal("test2")
	set.AddVal(1)
	set.AddVal(5.6)
	c := make(chan int)
	set.AddVal(c)
	if l := set.Len(); l != 5 {
		t.Error("Expected len=5, got len =", l)
	}

	// Delete all
	set.RemoveVal("test1")
	set.RemoveVal("test2")
	set.RemoveVal(1)
	set.RemoveVal(5.6)
	set.RemoveVal(c)
	if !set.IsEmpty() {
		t.Error("Expected len=0, got len =", set.Len())
	}
}

func TestHasVal(t *testing.T) {
	set := NewSet(5, 34.5, "testing")

	// Test table
	tests := []struct {
		input    interface{}
		expected bool
	}{
		{5, true},
		{34.5, true},
		{"testing", true},
		{"", false},
		{0, false},
		{0.0, false},
		{nil, false},
		{"hello", false},
		{false, false},
		{true, false},
	}

	for _, test := range tests {
		actual := set.HasVal(test.input)
		if actual != test.expected {
			t.Errorf("on set.HasVal(%v), expected %v, got %v",
				test.input, test.expected, actual)
		}
	}
}

func TestAddVal(t *testing.T) {
	set := NewSet()

	// Add One
	set.AddVal(5)
	l := set.Len()
	hasVal := set.HasVal(5)
	if l != 1 || !hasVal {
		t.Error("Expected set to contain only the int 5, got set =", set.set)
	}

	// Add several
	set.AddVal(45.6)
	set.AddVal("Hello")
	set.AddVal(42)
	set.AddVal("H")
	if !set.HasVal(45.6) || !set.HasVal("Hello") || !set.HasVal(42) || !set.HasVal("H") {
		t.Error("Add failed")
	}
}

func TestEquals(t *testing.T) {
	s1 := NewSet()
	s2 := NewSet()

	// Test empties
	if !s1.Equals(s2) {
		t.Error("sets should be equal; got:", s1, s2)
	}

	s1.AddVal(3)
	if s1.Equals(s2) {
		t.Error("sets should not be equal; got:", s1, s2)
	}

	s2.AddVal(3)
	if !s1.Equals(s2) {
		t.Error("sets should be equal; got:", s1, s2)
	}
}

func TestUnion(t *testing.T) {
	s1 := NewSet(1, 2, 3, 4, 5)
	s2 := NewSet(1.0, 2.0, 3.0)
	s3 := NewSet("a", "b", "c")
	s4 := NewSet(true, false)
	s5 := NewSet()

	// Test table
	tests := []struct {
		in       []*Set
		expected *Set
	}{
		{[]*Set{s1, s2}, NewSet(1, 2, 3, 4, 5, 1.0, 2.0, 3.0)},
		{[]*Set{s1, s3}, NewSet(1, 2, 3, 4, 5, "a", "b", "c")},
		{[]*Set{s3, s4}, NewSet("a", "b", "c", true, false)},
		{[]*Set{s3, s5}, s3},
		{[]*Set{s5, s3}, s3},
		{[]*Set{s5, s5}, s5},
	}

	for _, test := range tests {
		actual := test.in[0].Union(test.in[1])
		if !(test.expected.Equals(actual)) {
			t.Errorf("for union of sets %v and %v, expected %v, got %v",
				test.in[0], test.in[1], test.expected, actual)
		}
	}
}

func TestInter(t *testing.T) {
	s1 := NewSet(1, 2, 3, 4, 5, true)
	s2 := NewSet(1.0, 2.0, 3.0, "a")
	s3 := NewSet("a", "b", "c", 1, 2, 3)
	s4 := NewSet(true, false, 1, 1.0, "c")
	s5 := NewSet()

	// Test table
	tests := []struct {
		in       []*Set
		expected *Set
	}{
		{[]*Set{s1, s2}, s5},
		{[]*Set{s1, s5}, s5},
		{[]*Set{s5, s5}, s5},
		{[]*Set{s3, s1}, NewSet(1, 2, 3)},
		{[]*Set{s1, s3}, NewSet(1, 2, 3)},
		{[]*Set{s2, s3}, NewSet("a")},
		{[]*Set{s4, s1}, NewSet(true, 1)},
		{[]*Set{s2, s4}, NewSet(1.0)},
	}

	for _, test := range tests {
		actual := test.in[0].Inter(test.in[1])
		if !(test.expected.Equals(actual)) {
			t.Errorf("for intersection of sets %v and %v, expected %v, got %v",
				test.in[0], test.in[1], test.expected, actual)
		}
	}
}

func TestDiff(t *testing.T) {
	s1 := NewSet(1, 2, 3, 4, 5, true)
	s2 := NewSet(1.0, 2.0, 3.0, "a")
	s3 := NewSet("a", "b", "c", 1, 2, 3)
	s4 := NewSet(true, false, 1, 1.0, "c")
	s5 := NewSet()

	// Test table
	tests := []struct {
		in       []*Set
		expected *Set
	}{
		{[]*Set{s1, s2}, s1},
		{[]*Set{s1, s5}, s1},
		{[]*Set{s5, s5}, s5},
		{[]*Set{s3, s1}, NewSet("a", "b", "c")},
		{[]*Set{s1, s3}, NewSet(4, 5, true)},
		{[]*Set{s2, s3}, NewSet(1.0, 2.0, 3.0)},
		{[]*Set{s4, s1}, NewSet(false, 1.0, "c")},
		{[]*Set{s3, s4}, NewSet("a", "b", 2, 3)},
	}

	for _, test := range tests {
		actual := test.in[0].Diff(test.in[1])
		if !(test.expected.Equals(actual)) {
			t.Errorf("for difference of sets %v and %v, expected %v, got %v",
				test.in[0], test.in[1], test.expected, actual)
		}
	}
}

func TestPrint(_ *testing.T) {
	// Not sure how else to test the Set.String() function.
	// I can test unordered output on different lines, but not within the same line.
	// Output of this function is:
	//     Set( 1 2 3 4 5 6.7 "hello" "there" true false)
	// but the elements may be in any order.
	s := NewSet(1, 2, 2, 2, 3, 4, 5, 6.7, "hello", "there", true, false, "hello", "hello")
	fmt.Println(s)
}
