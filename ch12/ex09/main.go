// Copyright © 2017 Yuki Nagahara
// 練習12-9 ストリームトークンデコーダ

package main

import (
	"io"
	"log"
	"os"

	"github.com/naga718/golang-practice/ch12/ex09/sexpr"
)

func main() {
	d := sexpr.NewDecoder(os.Stdin)
	for {
		token, err := d.Token()
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}

		// トークンごとに分けて出力ができる
		switch token.(type) {
		case sexpr.StartList:
			t := token.(sexpr.StartList)
			log.Printf("[StartList] %v\n", string(t))
		case sexpr.EndList:
			t := token.(sexpr.EndList)
			log.Printf("[EndList] %v\n", string(t))
		case sexpr.Symbol:
			t := token.(sexpr.Symbol)
			log.Printf("[Symbol] %v\n", string(t))
		case sexpr.String:
			t := token.(sexpr.String)
			log.Printf("[String] %v\n", t)
		case sexpr.Int:
			t := token.(sexpr.Int)
			log.Printf("[Int] %v\n", t)
		}
	}
}
