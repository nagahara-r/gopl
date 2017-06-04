// Copyright © Yuki Nagahara

package tempconv

import "testing"

func TestKelvinSupport(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"0K",
			"-273.15°C",
		}, {
			"273.15K",
			"0°C",
		}, {
			"10000K",
			"9726.85°C",
		}, {
			"",
			"0°C",
		},
	}

	for _, test := range tests {
		cf := new(celsiusFlag)
		cf.Set(test.input)

		if cf.String() != test.expected {
			t.Errorf("input: %v, expected: %v Set(): %v\n", test.input, test.expected, cf.String())
		}
	}
}
