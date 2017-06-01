// Copyright © 2017 Yuki Nagahara
// 練習5-19: return文を含んでいないのに、ゼロ値ではない値を返す関数

package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Printf("returnTrueWithPanic() = %v\n", returnTrueWithPanic())
}

// returnTrueWithPanic は return文がないのにtrueを返します。
// (bool のゼロ値は false)
func returnTrueWithPanic() (b bool) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("%v\n", err)
			b = true
		}
	}()

	panic("panic!")
}
