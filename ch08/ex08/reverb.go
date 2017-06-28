// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習8-8: エコーサーバにタイムアウトを追加し、10秒以内に何も叫ばないクライアントとの接続を切るようにします。

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
	wg    = new(sync.WaitGroup)
	mutex = new(sync.Mutex)
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

func read(input *bufio.Scanner, inputch chan bool) {
	for input.Scan() {
		inputch <- true
	}

	inputch <- false
}

//!+
func handleConn(c *net.TCPConn) {
	input := bufio.NewScanner(c)
	inputch := make(chan bool)
	timeout := time.After(time.Second * 10)

	go read(input, inputch)

	for end := false; end == false; {
		select {
		case result := <-inputch:
			if !result {
				end = true
				break
			}
			go echo(c, input.Text(), 1*time.Second)
			timeout = time.After(time.Second * 10)

		case <-timeout:
			end = true
			break
		}
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
