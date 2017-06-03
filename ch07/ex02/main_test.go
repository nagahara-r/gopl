// Copyright © 2017 Yuki Nagahara

package main

import (
	"bytes"
	"testing"
)

func TestWordCounter(t *testing.T) {
	tests := []struct {
		input    []byte
		expected int64
	}{
		{
			[]byte("hello"),
			5,
		}, {
			[]byte("hello world"),
			11,
		}, {
			[]byte("1 2 3 4 5 6 7 8 9 10"),
			20,
		}, {
			[]byte(""),
			0,
		}, {
			nil,
			0,
		},
	}

	for _, test := range tests {
		bcw, result := CountingWriter(new(bytes.Buffer))
		bcw.Write(test.input)
		if *result != test.expected {
			t.Errorf("input = %v, expected = %v, Write() = %v", string(test.input), test.expected, *result)
		}
	}
}
