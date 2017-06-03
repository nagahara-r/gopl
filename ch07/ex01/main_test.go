// Copyright © 2017 Yuki Nagahara

package main

import "testing"

func TestWordCounter(t *testing.T) {
	tests := []struct {
		input    []byte
		expected int
	}{
		{
			[]byte("hello"),
			1,
		}, {
			[]byte("hello world\n"),
			2,
		}, {
			[]byte("1 2 3 4 5 6 7 8 9 10 words \n"),
			11,
		}, {
			[]byte(""),
			0,
		}, {
			nil,
			0,
		},
	}

	for _, test := range tests {
		var result WordCounter
		result.Write(test.input)
		if int(result) != test.expected {
			t.Errorf("input = %v, expected = %v, Write() = %v", string(test.input), test.expected, result)
		}
	}
}

func TestReturnCounter(t *testing.T) {
	tests := []struct {
		input    []byte
		expected int
	}{
		{
			[]byte("hello"),
			0,
		}, {
			[]byte("hello world\n"),
			1,
		}, {
			[]byte("1 2 3 4 5\r\n 6\r\n 7 8 9 10 words \r\n"),
			3,
		}, {
			[]byte(""),
			0,
		}, {
			nil,
			0,
		},
	}

	for _, test := range tests {
		var result ReturnCounter
		result.Write(test.input)
		if int(result) != test.expected {
			t.Errorf("input = %v, expected = %v, Write() = %v", string(test.input), test.expected, result)
		}
	}
}
