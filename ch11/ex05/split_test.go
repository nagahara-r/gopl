// Copyright © 2017 Yuki Nagahara
// 練習11-5: 入力と期待される表を使うように TestSplit を拡張します。

package split

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		s        string
		sep      string
		expected int
	}{
		{
			"a:b:c",
			":",
			3,
		}, {
			"あ、い、う、え、お",
			"、",
			5,
		}, {
			"aaa:",
			":",
			2,
		}, {
			"split every words",
			"",
			17,
		}, {
			"日本語もUTF-8でSplit",
			"",
			15,
		}, {
			"",
			":",
			1,
		}, {
			"",
			"",
			0,
		},
	}

	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		l := len(words)
		if l != test.expected {
			t.Errorf("Split(%v, %v) returned %v words, but expected %v words", test.s, test.sep, l, test.expected)
		}
	}
}
