package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	tests := []struct {
		input    uint64
		expected int
	}{
		{0, 0},
		{0x1, 1},
		{0x3, 2},
		{0x1234567890ABCDEF, 32},
		{0xFFFFFFFFFFFFFFFF, 64},
	}

	for _, test := range tests {
		result := PopCount(test.input)
		if result != test.expected {
			t.Errorf("PopCount(%v) = %v, want = %v", test.input, result, test.expected)
		}
	}
}

func TestPopCountByLoop(t *testing.T) {
	tests := []struct {
		input    uint64
		expected int
	}{
		{0, 0},
		{0x1, 1},
		{0x3, 2},
		{0x1234567890ABCDEF, 32},
		{0xFFFFFFFFFFFFFFFF, 64},
	}

	for _, test := range tests {
		result := PopCountByLoop(test.input)
		if result != test.expected {
			t.Errorf("PopCountByLoop(%v) = %v, want = %v", test.input, result, test.expected)
		}
	}
}

func TestPopCountByBitshift(t *testing.T) {
	tests := []struct {
		input    uint64
		expected int
	}{
		{0, 0},
		{0x1, 1},
		{0x3, 2},
		{0x1234567890ABCDEF, 32},
		{0xFFFFFFFFFFFFFFFF, 64},
	}

	for _, test := range tests {
		result := PopCountByBitshift(test.input)
		if result != test.expected {
			t.Errorf("PopCountByBitshift(%v) = %v, want = %v", test.input, result, test.expected)
		}
	}
}

func TestPopCountByBitclear(t *testing.T) {
	tests := []struct {
		input    uint64
		expected int
	}{
		{0, 0},
		{0x1, 1},
		{0x3, 2},
		{0x1234567890ABCDEF, 32},
		{0xFFFFFFFFFFFFFFFF, 64},
	}

	for _, test := range tests {
		result := PopCountByBitclear(test.input)
		if result != test.expected {
			t.Errorf("PopCountByBitclear(%v) = %v, want = %v", test.input, result, test.expected)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByLoop(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByBitshift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByBitshift(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByBitclear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByBitclear(0x1234567890ABCDEF)
	}
}

// $ sh testall.bash
// ok      command-line-arguments  0.007s
// BenchmarkPopCount-4                     2000000000               0.38 ns/op
// BenchmarkPopCountByLoop-4               50000000                26.4 ns/op
// BenchmarkPopCountByBitshift-4           20000000                97.9 ns/op
// BenchmarkPopCountByBitclear-4           50000000                25.1 ns/op
// 8回加算代入するLoopと 32回ビット演算するBitclearがあまり変わらない
