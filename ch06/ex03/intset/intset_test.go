// Copyright © 2017 Yuki Nagahara

package intset

import "testing"

func TestIntersectWith(t *testing.T) {
	tests := []struct {
		input    IntSet
		target   IntSet
		expected string
	}{
		{
			IntSet{[]uint64{3}},
			IntSet{[]uint64{2}},
			"{1}",
		}, {
			IntSet{[]uint64{2046}}, // 11111111110
			IntSet{[]uint64{555}},  //  1000101011
			"{1 3 5 9}",
		}, {
			IntSet{[]uint64{2047}}, //   11111111111
			IntSet{[]uint64{3000}}, //  101110111000
			"{3 4 5 7 8 9}",
		}, {
			IntSet{[]uint64{}},
			IntSet{[]uint64{3000}}, //  101110111000
			"{}",
		}, {
			IntSet{nil},
			IntSet{[]uint64{3000}}, //  101110111000
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
			IntSet{[]uint64{14}}, // 1110
			IntSet{[]uint64{7}},  //  111
			"{3}",
		}, {
			IntSet{[]uint64{2046}}, // 11111111110
			IntSet{[]uint64{555}},  //  1000101011
			"{2 4 6 7 8 10}",
		}, {
			IntSet{[]uint64{2047}}, //   11111111111
			IntSet{[]uint64{3000}}, //  101110111000
			"{0 1 2 6 10}",
		}, {
			IntSet{[]uint64{}},
			IntSet{[]uint64{3000}}, //  101110111000
			"{}",
		}, {
			IntSet{nil},
			IntSet{[]uint64{3000}}, //  101110111000
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
			IntSet{[]uint64{15}}, // 1111
			IntSet{[]uint64{7}},  //  111
			"{3}",
		}, {
			IntSet{[]uint64{2046}}, // 11111111110
			IntSet{[]uint64{555}},  //  1000101011
			"{0 2 4 6 7 8 10}",
		}, {
			IntSet{[]uint64{2047}}, //   11111111111
			IntSet{[]uint64{3000}}, //  101110111000
			"{0 1 2 6 10 11}",
		}, {
			IntSet{[]uint64{}},
			IntSet{[]uint64{2047}}, //  11111111111
			"{0 1 2 3 4 5 6 7 8 9 10}",
		}, {
			IntSet{nil},
			IntSet{[]uint64{3000}}, //  101110111000
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
