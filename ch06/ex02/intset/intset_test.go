// Copyright © 2017 Yuki Nagahara

package intset

import "testing"

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
