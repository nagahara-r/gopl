// Copyright © 2017 Yuki Nagahara

package treesort

import "testing"

func TestString(t *testing.T) {
	tests := []struct {
		input    []int
		expected string
	}{
		{
			[]int{0, 1, 2, 3},
			"[0 1 2 3]",
		}, {
			[]int{3, 5, 1, 4, 2},
			"[1 2 3 4 5]",
		}, {
			[]int{100, 0, -5, 18, 42, 33, 99, 45, 36, 45},
			"[-5 0 18 33 36 42 45 45 99 100]",
		}, {
			nil,
			"[]",
		},
	}

	for _, test := range tests {
		var tr *tree
		for _, val := range test.input {
			tr = add(tr, val)
		}

		result := tr.String()
		if test.expected != result {
			t.Errorf("input = %v, expected = %v, String() = %v", test.input, test.expected, result)
		}
	}
}
