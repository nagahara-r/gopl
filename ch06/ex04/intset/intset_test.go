// Copyright © 2017 Yuki Nagahara

package intset

import "testing"

func TestElems(t *testing.T) {
	tests := []struct {
		input    IntSet
		expected []int
	}{
		{
			IntSet{[]uint64{1}}, // 1
			[]int{0},
		}, {
			IntSet{[]uint64{2046}}, // 11111111110
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		}, {
			IntSet{[]uint64{2047}}, //   11111111111
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		}, {
			IntSet{[]uint64{12345}}, // 11000000111001
			[]int{0, 3, 4, 5, 12, 13},
		}, {
			IntSet{[]uint64{0, 1, 0, 0, 0}},
			[]int{64},
		}, {
			IntSet{nil},
			[]int{},
		},
	}

	for _, test := range tests {
		if !comp(test.expected, test.input.Elems()) {
			t.Errorf("expected = %v, Elems() = %v", test.expected, test.input.Elems())
		}
	}
}

func comp(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
