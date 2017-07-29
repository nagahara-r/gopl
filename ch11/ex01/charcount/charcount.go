// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習11-1: Charcount のテストを作成

// See page 97.
//!+

// Package charcount computes counts of Unicode characters.
package charcount

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

// CharCount はReaderに与えられた入力の文字数カウントを実施します。
// テストできるように関数を切り離しています
func CharCount(r io.Reader) (counts map[rune]int, utflen [utf8.UTFMax + 1]int, invalid int) {
	if r == nil {
		return
	}
	counts = make(map[rune]int) // counts of Unicode characters

	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}

	return
}

//!-
