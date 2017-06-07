// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習7-5：LimitReader の実装

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

// LReader は、String用のio.Readerです。
type LReader struct {
	r io.Reader
	n int64
}

func (lr *LReader) Read(p []byte) (n int, err error) {
	if lr.n <= 0 {
		return 0, io.EOF
	}

	n = len(p)

	if int64(n) > lr.n {
		n = int(lr.n)
	}

	n, err = lr.r.Read(p[:n])
	if err != nil {
		return n, err
	}

	lr.n -= int64(n)
	return n, err
}

// LimitReader は指定されたバイト数まで読み出します。
func LimitReader(r io.Reader, n int64) io.Reader {
	return &LReader{r, n}
}

func main() {
	r := strings.NewReader(string("1234567890"))

	lr := LimitReader(r, 5)

	buf, err := ioutil.ReadAll(lr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read = %v\n", string(buf))
}
