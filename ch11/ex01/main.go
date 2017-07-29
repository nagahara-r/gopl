// Copyright © 2017 Yuki Nagahara

package main

import (
	"fmt"
	"os"

	"github.com/naga718/golang-practice/ch11/ex01/charcount"
)

func main() {
	counts, utflen, invalid := charcount.CharCount(os.Stdin)

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
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
