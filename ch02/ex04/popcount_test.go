package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	PopCount(0x1234567890ABCDEF)
}

func TestPopCountByLoop(t *testing.T) {
	PopCountByLoop(0x1234567890ABCDEF)
}

func TestPopCountByBitshift(t *testing.T) {
	PopCountByBitshift(0x1234567890ABCDEF)
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

// $ sh testall.bash
// ok      command-line-arguments  0.006s
// BenchmarkPopCount-4                     2000000000               0.40 ns/op
// BenchmarkPopCountByLoop-4               50000000                26.4 ns/op
// BenchmarkPopCountByBitshift-4           20000000                98.4 ns/op
// Bitshiftはさらに演算回数、ビットシフト回数が増加
