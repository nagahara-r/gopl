// Copyright © 2017 Yuki Nagahara

package equal

import (
	"math"
	"testing"
)

func TestFloat(t *testing.T) {
	tests := []struct {
		x        float64
		y        float64
		expected bool
	}{
		{float64(1) / float64(3), float64(0.333333333), true},
		{1 + float64(0.333333333), float64(0.333333333), false},
		{float64(0.3333333332), float64(0.3333333339), true},
		{math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64 + float64(0.0000000001), true},
	}

	for _, test := range tests {
		result := Equal(test.x, test.y)

		if result != test.expected {
			t.Errorf("Equal(%f, %f) = %v, but expected %v", test.x, test.y, result, test.expected)
		}
	}
}

func TestComplex(t *testing.T) {
	tests := []struct {
		x        complex128
		y        complex128
		expected bool
	}{
		{complex(float64(1)/float64(3), 1), complex(0.333333333, 1), true},
		{complex(1+float64(0.333333333), 3), complex(float64(0.333333333), 3), false},
		{complex(float64(0.3333333332), 10), complex(float64(0.3333333339), 10), true},
		{complex(math.SmallestNonzeroFloat64, 1), complex(math.SmallestNonzeroFloat64+float64(0.0000000001), 1), true},
		{complex(float64(0.3333333332), 10), complex(float64(0.3333333339), 0), false},
	}

	for _, test := range tests {
		result := Equal(test.x, test.y)

		if result != test.expected {
			t.Errorf("Equal(%f, %f) = %v, but expected %v", test.x, test.y, result, test.expected)
		}
	}
}
