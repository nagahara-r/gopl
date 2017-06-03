// Copyright © 2017 Yuki Nagahara

package intset

import "testing"

func TestElems(t *testing.T) {
	tests := []struct {
		input    IntSet
		expected []int
	}{
		{
			IntSet{[]uint{1}}, // 1
			[]int{0},
		}, {
			IntSet{[]uint{2046}}, // 11111111110
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		}, {
			IntSet{[]uint{2047}}, //   11111111111
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		}, {
			IntSet{[]uint{12345}}, // 11000000111001
			[]int{0, 3, 4, 5, 12, 13},
		}, {
			IntSet{[]uint{0, 1, 0, 0, 0}},
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

func TestIntersectWith(t *testing.T) {
	tests := []struct {
		input    IntSet
		target   IntSet
		expected string
	}{
		{
			IntSet{[]uint{3}},
			IntSet{[]uint{2}},
			"{1}",
		}, {
			IntSet{[]uint{2046}}, // 11111111110
			IntSet{[]uint{555}},  //  1000101011
			"{1 3 5 9}",
		}, {
			IntSet{[]uint{2047}}, //   11111111111
			IntSet{[]uint{3000}}, //  101110111000
			"{3 4 5 7 8 9}",
		}, {
			IntSet{[]uint{}},
			IntSet{[]uint{3000}}, //  101110111000
			"{}",
		}, {
			IntSet{nil},
			IntSet{[]uint{3000}}, //  101110111000
			"{}",
		},
	}

	for _, test := range tests {
		test.input.IntersectWith(&test.target)
		if test.expected != test.input.String() {
			t.Errorf("expected = %v, IntersectWith() = %v", test.expected, test.input.String())
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	tests := []struct {
		input    IntSet
		target   IntSet
		expected string
	}{
		{
			IntSet{[]uint{14}}, // 1110
			IntSet{[]uint{7}},  //  111
			"{3}",
		}, {
			IntSet{[]uint{2046}}, // 11111111110
			IntSet{[]uint{555}},  //  1000101011
			"{2 4 6 7 8 10}",
		}, {
			IntSet{[]uint{2047}}, //   11111111111
			IntSet{[]uint{3000}}, //  101110111000
			"{0 1 2 6 10}",
		}, {
			IntSet{[]uint{}},
			IntSet{[]uint{3000}}, //  101110111000
			"{}",
		}, {
			IntSet{nil},
			IntSet{[]uint{3000}}, //  101110111000
			"{}",
		},
	}

	for _, test := range tests {
		test.input.DifferenceWith(&test.target)
		if test.expected != test.input.String() {
			t.Errorf("expected = %v, DifferenceWith() = %v", test.expected, test.input.String())
		}
	}
}

func TestSymmetricDifference(t *testing.T) {
	tests := []struct {
		input    IntSet
		target   IntSet
		expected string
	}{
		{
			IntSet{[]uint{15}}, // 1111
			IntSet{[]uint{7}},  //  111
			"{3}",
		}, {
			IntSet{[]uint{2046}}, // 11111111110
			IntSet{[]uint{555}},  //  1000101011
			"{0 2 4 6 7 8 10}",
		}, {
			IntSet{[]uint{2047}}, //   11111111111
			IntSet{[]uint{3000}}, //  101110111000
			"{0 1 2 6 10 11}",
		}, {
			IntSet{[]uint{}},
			IntSet{[]uint{2047}}, //  11111111111
			"{0 1 2 3 4 5 6 7 8 9 10}",
		}, {
			IntSet{nil},
			IntSet{[]uint{3000}}, //  101110111000
			"{3 4 5 7 8 9 11}",
		},
	}

	for _, test := range tests {
		test.input.SymmetricDifference(&test.target)
		if test.expected != test.input.String() {
			t.Errorf("expected = %v, SymmetricDifference() = %v", test.expected, test.input.String())
		}
	}
}

func TestAddAll(t *testing.T) {
	tests := []struct {
		input    []int
		expected string
	}{
		{
			[]int{0, 1, 2, 3},
			"{0 1 2 3}",
		}, {
			[]int{8, 3, 1, 2},
			"{1 2 3 8}",
		}, {
			[]int{},
			"{}",
		}, {
			nil,
			"{}",
		},
	}

	for _, test := range tests {
		result := IntSet{}
		result.AddAll(test.input...)
		if test.expected != result.String() {
			t.Errorf("expected = %v, AddAll() = %v", test.expected, result.String())
		}
	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		input    IntSet
		expected int
	}{
		{
			IntSet{[]uint{15}},
			4,
		}, {
			IntSet{[]uint{8}},
			1,
		}, {
			IntSet{[]uint{1, 3}},
			3,
		}, {
			IntSet{[]uint{}},
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
			IntSet{[]uint{2}},
			1,
			IntSet{[]uint{0}},
		}, {
			IntSet{[]uint{128}},
			7,
			IntSet{[]uint{0}},
		}, {
			IntSet{[]uint{1, 1}},
			64,
			IntSet{[]uint{1, 0}},
		}, {
			IntSet{[]uint{0, 1}},
			-1,
			IntSet{[]uint{0, 1}},
		}, {
			IntSet{[]uint{1}},
			0,
			IntSet{[]uint{0}},
		}, {
			IntSet{[]uint{12345}},
			65,
			IntSet{[]uint{12345}},
		}, {
			IntSet{nil},
			0,
			IntSet{nil},
		},
	}

	for _, test := range tests {
		test.input.Remove(test.remove)
		if !compIntSet(test.input, test.expected) {
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
			IntSet{[]uint{1}},
			IntSet{[]uint{}},
		}, {
			IntSet{[]uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
			IntSet{[]uint{}},
		}, {
			IntSet{[]uint{}},
			IntSet{[]uint{}},
		}, {
			IntSet{nil},
			IntSet{[]uint{}},
		},
	}

	for _, test := range tests {
		test.input.Clear()
		if !compIntSet(test.input, test.expected) {
			t.Errorf("expected = %v, Clear() = %v", test.expected, test.input)
		}
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		input IntSet
	}{
		{
			IntSet{[]uint{0, 1, 2, 3}},
		}, {
			IntSet{[]uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		}, {
			IntSet{[]uint{}},
		}, {
			IntSet{nil},
		},
	}

	for _, test := range tests {
		result := test.input.Copy()
		if !compIntSet(test.input, *result) {
			t.Errorf("expected = %v, Copy() = %v", test.input, test.input)
		}
	}
}

func compIntSet(a IntSet, b IntSet) bool {
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
