package main

// Copyright © 2017 Yuki Nagahara
// 練習9-5: メッセージを送り合う2つのゴルーチンを作成します。
// タイマー付きで1秒間に送り合うメッセージの数を確認します。

import (
	"flag"
	"fmt"
	"time"
)

var pingpong = make(chan struct{})
var rally int

func main() {
	t := flag.Int64("t", 1, "Rally time (Sec)")
	flag.Parse()

	timeout := time.After(time.Second * time.Duration(*t))

	go ping()
	go pong()

	select {
	case <-timeout:
		fmt.Printf("%v Times\n", rally)
	}
}

func ping() {
	for {
		rally++
		pingpong <- struct{}{}
		<-pingpong
	}
}

func pong() {
	for {
		<-pingpong
		pingpong <- struct{}{}
	}
}
