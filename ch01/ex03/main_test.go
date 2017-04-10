package main

import (
	"os"
	"testing"
)

func TestEchoWithIndexLoop(t *testing.T) {
	os.Args = []string{"main.go", "test"}
	echoWithIndexLoop()
}

func TestEchoWithRangeLoop(t *testing.T) {
	os.Args = []string{"main.go", "test"}
	echoWithRangeLoop()
}

func TestEchoWithJoin(t *testing.T) {
	os.Args = []string{"main.go", "test"}
	echoWithJoin()
}

func BenchmarkEchoWithIndexLoop(b *testing.B) {
	os.Args = []string{"main.go", "test"}
	for i := 0; i < b.N; i++ {
		echoWithIndexLoop()
	}
}

func BenchmarkEchoWithRangeLoop(b *testing.B) {
	os.Args = []string{"main.go", "test"}
	for i := 0; i < b.N; i++ {
		echoWithRangeLoop()
	}
}

func BenchmarkEchoWithJoin(b *testing.B) {
	os.Args = []string{"main.go", "test"}
	for i := 0; i < b.N; i++ {
		echoWithJoin()
	}
}

// $ sh testall.bash
// BenchmarkEchoWithIndexLoop-4    100000000               15.2 ns/op
// BenchmarkEchoWithRangeLoop-4    100000000               15.4 ns/op
// BenchmarkEchoWithJoin-4         300000000                5.75 ns/op
// PASS
// ok      github.com/naga718/golang-practice/ch01/ex03    5.404s
