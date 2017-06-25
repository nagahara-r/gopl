// Copyright © 2017 Yuki Nagahara
// 練習8-4: Reverb2 を修正し、echoゴルーチンが動作中は待つようにします。

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"testing"
)

var (
	stdin  io.Reader = os.Stdin
	stdout io.Writer = os.Stdout
)

// netcat3 は練習8-3相当のプログラムです。
// Stdin Stdoutを差し替えできるようにのみ変更しています
func netcat3() {
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
		n, _ := io.Copy(stdout, tcpconn) // NOTE: ignoring errors

		log.Printf("read = %v", n)
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()

	mustCopy(tcpconn, stdin)
	tcpconn.CloseWrite()
	<-done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	var n int64
	var err error
	if n, err = io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
	log.Printf("write = %v", n)
}

// 正しくすべてのechoが返ってくるかテストします
func TestEcho(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			"Aa\n",
			[]string{strings.ToUpper("Aa"), "Aa", strings.ToLower("Aa")},
		},
		{
			"Aa\nBb\nCc\n",
			[]string{strings.ToUpper("Aa"), "Aa", strings.ToLower("Aa"), strings.ToUpper("Bb"), "Bb", strings.ToLower("Bb"), strings.ToUpper("Cc"), "Cc", strings.ToLower("Cc")},
		},
	}

	go main()

	for _, test := range tests {
		buf := new(bytes.Buffer)
		stdin = bytes.NewBufferString(test.input)
		stdout = buf

		netcat3()
		if !verify(string(buf.Bytes()), test.expected) {
			t.Errorf("get = %v, expected echos = %v", string(buf.Bytes()), test.expected)
		}
	}
}

func verify(get string, strs []string) (ok bool) {
	for _, str := range strs {
		if !strings.Contains(get, str) {
			return false
		}
	}

	return true
}
