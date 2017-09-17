// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習13-3: bzip圧縮を並行に動作できるようにします。

// See page 365.

// Bzipper reads input, bzip2-compresses it, and writes it out.
package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/naga718/golang-practice/ch13/ex03/bzip"
)

var (
	wg = new(sync.WaitGroup)
)

func main() {
	for i, arg := range os.Args[1:] {
		r, err := os.Open(arg)
		if err != nil {
			log.Fatal(err)
		}

		w, err := os.Create("out/" + strconv.Itoa(i) + ".bin")
		if err != nil {
			log.Fatal(err)
		}
		worker(r, w)
	}
	wg.Wait()
}

func worker(r io.Reader, w io.Writer) {
	wg.Add(1)
	go func(r io.Reader, w io.Writer) {
		defer wg.Done()
		bzw := bzip.NewWriter(w)
		if _, err := io.Copy(bzw, r); err != nil {
			log.Fatalf("bzipper: %v\n", err)
		}
		if err := bzw.Close(); err != nil {
			log.Fatalf("bzipper: close: %v\n", err)
		}
	}(r, w)
}
