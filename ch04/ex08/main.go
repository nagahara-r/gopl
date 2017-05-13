// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Copyright © 2017 Yuki Nagahara
// CharCount をUnicode分類に従って文字数、数字数などを数えます。

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type unicodeKindCount struct {
	control int
	digit   int
	graphic int
	letter  int
	lower   int
	mark    int
	number  int
	print   int
	punct   int
	space   int
	symbol  int
	title   int
	upper   int
}

func printUnicodeKindCount(u unicodeKindCount) {
	fmt.Printf("Unicode Kind Count\n")
	fmt.Printf("Control\t%d\n", u.control)
	fmt.Printf("Digit\t%d\n", u.digit)
	fmt.Printf("Graphic\t%d\n", u.graphic)
	fmt.Printf("Letter\t%d\n", u.letter)
	fmt.Printf("Lower\t%d\n", u.lower)
	fmt.Printf("Mark\t%d\n", u.mark)
	fmt.Printf("Number\t%d\n", u.number)
	fmt.Printf("Print\t%d\n", u.print)
	fmt.Printf("Punct\t%d\n", u.punct)
	fmt.Printf("Space\t%d\n", u.space)
	fmt.Printf("Symbol\t%d\n", u.symbol)
	fmt.Printf("Title\t%d\n", u.title)
	fmt.Printf("Upper\t%d\n", u.upper)
}

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	var ucount unicodeKindCount

	in := bufio.NewReader(os.Stdin)
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
		if unicode.IsControl(r) {
			ucount.control++
		}
		if unicode.IsDigit(r) {
			ucount.digit++
		}
		if unicode.IsGraphic(r) {
			ucount.graphic++
		}
		if unicode.IsLetter(r) {
			ucount.letter++
		}
		if unicode.IsLower(r) {
			ucount.lower++
		}
		if unicode.IsMark(r) {
			ucount.mark++
		}
		if unicode.IsNumber(r) {
			ucount.number++
		}
		if unicode.IsPrint(r) {
			ucount.print++
		}
		if unicode.IsPunct(r) {
			ucount.punct++
		}
		if unicode.IsSpace(r) {
			ucount.space++
		}
		if unicode.IsSymbol(r) {
			ucount.symbol++
		}
		if unicode.IsTitle(r) {
			ucount.title++
		}
		if unicode.IsUpper(r) {
			ucount.upper++
		}

		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	fmt.Print("\n")
	printUnicodeKindCount(ucount)
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//!-
