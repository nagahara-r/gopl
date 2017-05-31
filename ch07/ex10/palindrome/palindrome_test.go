package palindrome

import (
	"sort"
	"testing"
)

// Copyright © 2017 Yuki Nagahara

func TestInt(t *testing.T) {
	var tests = []struct {
		input    []int
		expected bool
	}{
		{
			[]int{0, 1, 2, 3, 2, 1, 0},
			true,
		},
		{
			[]int{0, 1, 1, 1, 0},
			true,
		},
		{
			[]int{5, 5, 5, 5},
			true,
		},
		{
			[]int{0, 0},
			true,
		},
		{
			[]int{0},
			true,
		},
		{
			[]int{},
			true,
		},
		{
			[]int{0, 1, 2, 3, 4},
			false,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 8, 9},
			false,
		},
		{
			[]int{0, 1},
			false,
		},
	}

	for _, test := range tests {
		result := IsPalindrome(sort.IntSlice(test.input))
		if result != test.expected {
			t.Errorf("input: %v, expected: %v, result: %v\n", test.input, test.expected, result)
		}
	}
}
