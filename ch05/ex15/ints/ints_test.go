package ints

import "testing"

// Copyright © 2017 Yuki Nagahara

func TestMax(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{2, 4, 3, 0, 1},
			4,
		},
		{
			[]int{-1, 0, 1},
			1,
		},
		{
			[]int{-1},
			-1,
		},
		{
			[]int{},
			0,
		},
		{
			nil,
			0,
		},
	}

	for _, test := range tests {
		result := Max(test.input...)
		if result != test.expected {
			t.Errorf("expected: %v, result: %v", test.expected, result)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{5, 4, 0, 2, 1},
			0,
		},
		{
			[]int{-1, 0, 1},
			-1,
		},
		{
			[]int{-1},
			-1,
		},
		{
			[]int{},
			0,
		},
		{
			nil,
			0,
		},
	}

	for _, test := range tests {
		result := Min(test.input...)
		if result != test.expected {
			t.Errorf("expected: %v, result: %v", test.expected, result)
		}
	}
}
