// Copyright © 2017 Yuki Nagahara

package strings

import "testing"

func TestStrings(t *testing.T) {
	tests := []struct {
		input    []string
		sep      string
		expected string
	}{
		{
			[]string{"a", "b", "c", "d"},
			"&",
			"a&b&c&d",
		},
		{
			[]string{"123", "456", "7890", "12"},
			"+",
			"123+456+7890+12",
		},
		{
			[]string{"日本語", "つなげること", "できる？"},
			"は",
			"日本語はつなげることはできる？",
		},
		{
			[]string{"abcde", "", "", "fgh"},
			".",
			"abcde...fgh",
		},
		{
			nil,
			"&",
			"",
		},
	}

	for _, test := range tests {
		result := Join(test.sep, test.input...)
		if test.expected != result {
			t.Errorf("expected = %v, Join() = %v", test.expected, result)
		}
	}
}
