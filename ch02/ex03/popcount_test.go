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

// $ go test -bench=.
// BenchmarkPopCount-4             2000000000               0.36 ns/op
// BenchmarkPopCountByLoop-4       50000000                26.4 ns/op
// PASS
// ok      github.com/naga718/golang-practice/ch02/ex03    2.113s
// メモリに保存する分演算のみのPopCountのほうが早い
