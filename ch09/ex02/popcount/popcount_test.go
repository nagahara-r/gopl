package popcount

import "testing"

func TestPopCount(t *testing.T) {
	tests := []struct {
		value    uint64
		expected int
	}{
		{1, 1},
		{3, 2},
		{67108863, 26},            // 11111111111111111111111111
		{9223372036854775807, 63}, // MAX of int64
		{0, 0},
	}

	for _, test := range tests {
		result := PopCount(test.value)
		if result != test.expected {
			t.Errorf("PopCount(%v) = %v, expected = %v", test.value, result, test.expected)
		}
	}
}
