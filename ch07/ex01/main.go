// Copyright © 2017 Yuki Nagahara
// 練習7-1 単語数、行数を数えるカウンタの実装

package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// WordCounter は単語数を数えます。
type WordCounter int

// ReturnCounter は行数を数えます。
type ReturnCounter int

func (wc *WordCounter) Write(str []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(str))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		*wc++
		//log.Printf("%v: %v", count, scanner.Text())
	}

	return int(*wc), nil
}

func (rc *ReturnCounter) Write(str []byte) (int, error) {
	runes := []rune(string(str))

	for _, r := range runes {
		if r == '\n' {
			*rc++
		}
	}

	return int(*rc), nil
}

func main() {
	//!+main
	var wc WordCounter
	var rc ReturnCounter
	wc.Write([]byte("hello"))
	rc.Write([]byte("hello"))
	fmt.Printf("words = %v, returns = %v\n", wc, rc)

	wc = 0 // reset counter
	rc = 0
	var name = "Golang"
	fmt.Fprintf(&wc, "hello, %s World!\n", name)
	fmt.Fprintf(&rc, "hello, %s World!\n", name)
	fmt.Printf("words = %v, returns = %v\n", wc, rc)
}
