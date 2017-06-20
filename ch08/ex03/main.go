// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習8-3: 標準入力が閉じられたあとでもエコーが終わるまで出力し続けるように修正

// See page 227.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

//!+
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	defer conn.Close()

	tcpconn, ok := conn.(*net.TCPConn)
	if !ok {
		err = fmt.Errorf("Dial TCPConn Assertion Error")
	}

	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, tcpconn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(tcpconn, os.Stdin)
	tcpconn.CloseWrite()
	<-done // wait for background goroutine to finish
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
