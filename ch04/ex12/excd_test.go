// Excd はウェブコミックExcdを検索します。
// Copyright © 2017 Yuki Nagahara

package main

import "testing"

func TestParseAll(t *testing.T) {
	want0 := "Barrel - Part 1"
	want10 := "Barrel - Part 2"
	excds := ParseAll()

	if excds[0].Title != want0 {
		t.Errorf("want = %v, get = %v", want0, excds[0].Title)
	}

	if excds[10].Title != want10 {
		t.Errorf("want = %v, get = %v", want10, excds[10].Title)
	}
}

func TestSearch(t *testing.T) {
	want := 2
	excds := ParseAll()

	result := Search(excds[:100], "Apple")

	if len(result) != want {
		t.Errorf("want = %v, get = %v", want, len(result))
	}
}
