// Copyright © 2017 Yuki Nagahara

package intset

import "testing"

func TestLen(t *testing.T) {
	tests := []struct {
		input    IntSet
		expected int
	}{
		{
			IntSet{[]uint64{15}},
			4,
		}, {
			IntSet{[]uint64{8}},
			1,
		}, {
			IntSet{[]uint64{1, 3}},
			3,
		}, {
			IntSet{[]uint64{}},
			0,
		}, {
			IntSet{nil},
			0,
		},
	}

	for _, test := range tests {
		result := test.input.Len()
		if test.expected != result {
			t.Errorf("expected = %v, Len() = %v", test.expected, result)
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		input    IntSet
		remove   int
		expected IntSet
	}{
		{
			IntSet{[]uint64{2}},
			1,
			IntSet{[]uint64{0}},
		}, {
			IntSet{[]uint64{128}},
			7,
			IntSet{[]uint64{0}},
		}, {
			IntSet{[]uint64{1, 1}},
			64,
			IntSet{[]uint64{1, 0}},
		}, {
			IntSet{[]uint64{0, 1}},
			-1,
			IntSet{[]uint64{0, 1}},
		}, {
			IntSet{[]uint64{1}},
			0,
			IntSet{[]uint64{0}},
		}, {
			IntSet{[]uint64{12345}},
			65,
			IntSet{[]uint64{12345}},
		}, {
			IntSet{nil},
			0,
			IntSet{nil},
		},
	}

	for _, test := range tests {
		test.input.Remove(test.remove)
		if !comp(test.input, test.expected) {
			t.Errorf("expected = %v, Remove() = %v", test.expected, test.input)
		}
	}
}

func TestClear(t *testing.T) {
	tests := []struct {
		input    IntSet
		expected IntSet
	}{
		{
			IntSet{[]uint64{1}},
			IntSet{[]uint64{}},
		}, {
			IntSet{[]uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
			IntSet{[]uint64{}},
		}, {
			IntSet{[]uint64{}},
			IntSet{[]uint64{}},
		}, {
			IntSet{nil},
			IntSet{[]uint64{}},
		},
	}

	for _, test := range tests {
		test.input.Clear()
		if !comp(test.input, test.expected) {
			t.Errorf("expected = %v, Clear() = %v", test.expected, test.input)
		}
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		input IntSet
	}{
		{
			IntSet{[]uint64{0, 1, 2, 3}},
		}, {
			IntSet{[]uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		}, {
			IntSet{[]uint64{}},
		}, {
			IntSet{nil},
		},
	}

	for _, test := range tests {
		result := test.input.Copy()
		if !comp(test.input, *result) {
			t.Errorf("expected = %v, Copy() = %v", test.input, test.input)
		}
	}
}

func comp(a IntSet, b IntSet) bool {
	if len(a.words) != len(b.words) {
		return false
	}

	for i := range a.words {
		if a.words[i] != b.words[i] {
			return false
		}
	}

	return true
}
