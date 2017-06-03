package main

// Copyright © 2017 Yuki Nagahara

import "testing"

func TestExpand(t *testing.T) {
	var tests = []struct {
		text     string
		f        func(string) string
		expected string
	}{
		{
			"expands $foo <- string!",
			func(a string) string { return a + "bar" },
			"expands foobar <- string!",
		},
		{
			"expands $foo $bar $hoge $piyo <- string!",
			func(a string) string { return a + a },
			"expands foofoo barbar hogehoge piyopiyo <- string!",
		},
		{
			"expands foo bar hoge piyo <- string!",
			func(a string) string { return a + "nop" },
			"expands foo bar hoge piyo <- string!",
		},
		{
			"expands $foo $$bar $$$hoge piyo <- string!",
			func(a string) string { return a + "!" },
			"expands foo! $bar! $$hoge! piyo <- string!",
		},
		{
			"expands $日本語 $日本語二つ目 <- string!",
			func(a string) string { return a + "です" },
			"expands 日本語です 日本語二つ目です <- string!",
		},
		{
			"$",
			func(a string) string { return a + "aaa" },
			"aaa",
		},
		{
			"expands $foo <- string!",
			nil,
			"expands $foo <- string!",
		},
	}

	for _, test := range tests {
		result := expand(test.text, test.f)
		if result != test.expected {
			t.Errorf("text: %v, expected: %v, result: %v", test.text, test.expected, result)
		}
	}
}
