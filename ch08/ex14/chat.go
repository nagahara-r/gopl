// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習8-14: 個々のクライアントが到着した際、名前を提供するように変更します。

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

//!+broadcaster
type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages

	clients = []string{}
)

const (
	timeoutsec = 300
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	//who := conn.RemoteAddr().String()
	input := bufio.NewScanner(conn)
	who := enterName(input, ch)

	clients = append(clients, who)

	ch <- "You are " + who
	ch <- "Clients : \n" + strClients(clients) + "\n"

	messages <- who + " has arrived"
	entering <- ch

	scanning := make(chan string)

	timeout := time.After(time.Second * timeoutsec)
	fin := make(chan struct{})
	go func(conn net.Conn) {
		for input.Scan() {
			scanning <- input.Text()
		}

		close(fin)

		// NOTE: ignoring potential errors from input.Err()
	}(conn)

	// Scanもゴルーチン化して、timeoutしない限りは新たなタイムアウト時間をセットし続ける
	for isFinish := false; isFinish == false; {
		select {
		case str := <-scanning:
			timeout = time.After(time.Second * timeoutsec)
			messages <- who + ": " + str
		case <-timeout:
			isFinish = true
		case <-fin:
			isFinish = true
		}
	}

	leaving <- ch
	messages <- who + " has left"
	clients = remove(clients, who)
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func strClients(srcs []string) (result string) {
	return strings.Join(srcs, "\n")
}

func remove(srcs []string, dst string) (result []string) {
	for _, src := range srcs {
		if src != dst {
			result = append(result, src)
		}
	}
	return result
}

func contains(srcs []string, target string) bool {
	for _, src := range srcs {
		if src == target {
			return true
		}
	}
	return false
}

func enterName(input *bufio.Scanner, ch chan string) (name string) {
	ch <- "Input Your Name"
	for input.Scan() {
		if contains(clients, input.Text()) {
			ch <- "Sorry! This Name is Used"
			continue
		}
		break
	}

	return input.Text()
}

//!-main
