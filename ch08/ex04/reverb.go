// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習8-4: Reverb2 を修正し、echoゴルーチンが動作中は待つようにします。

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var (
	wg = new(sync.WaitGroup)
)

func echo(c net.Conn, shout string, delay time.Duration) {
	wg.Add(1)
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	wg.Done()
}

//!+
func handleConn(c *net.TCPConn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}

	c.CloseRead()
	wg.Wait()

	c.CloseWrite()
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		log.Printf("Accept")
		tcpconn, ok := conn.(*net.TCPConn)
		if !ok {
			err = fmt.Errorf("Dial TCPConn Assertion Error")
		}
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}

		go handleConn(tcpconn)
	}
}
