// Copyright © 2017 Yuki Nagahara
// 練習9-3 Func および(*Memo).Get を拡張して、メソッドをキャンセルできるようにする。
package main

import (
	"bufio"
	"os"

	"github.com/naga718/golang-practice/ch09/ex03/memo"
	"github.com/naga718/golang-practice/ch09/ex03/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func main() {
	done := make(chan struct{})
	m := memo.New(httpGetBody)
	defer m.Close()

	go func() {
		sc := bufio.NewScanner(os.Stdin)
		//for {
		for sc.Scan() {
		}
		done <- struct{}{}
		//}
	}()
	memotest.Sequential(nil, m, done)
}
