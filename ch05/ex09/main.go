package main

// Copyright © 2017 Yuki Nagahara
// 練習5-9: $foo を f("foo") が返すテキストで置換する関数

import (
	"fmt"
	"strings"
)

func main() {
	f := func(a string) string {
		return a + "bar"
	}

	s := "expands $foo <- string!"
	fmt.Printf("before: %v\n", s)

	fmt.Printf("after: %v\n", expand(s, f))
}

// expand は $foo を f("foo") が返すテキストで置換します。
func expand(s string, f func(string) string) string {
	if f == nil {
		return s
	}

	sstring := strings.Split(s, " ")

	for i := range sstring {
		if strings.HasPrefix(sstring[i], "$") {
			sstring[i] = f(sstring[i][1:])
		}
	}

	return strings.Join(sstring, " ")
}
