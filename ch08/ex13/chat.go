// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習8-13: 5分間何も送ってこないクライアントを切断するように修正

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

	who := conn.RemoteAddr().String()
	clients = append(clients, who)

	ch <- "You are " + who
	ch <- "Clients : \n" + strClients(clients) + "\n"
	messages <- who + " has arrived"
	entering <- ch

	scanning := make(chan string)
	go func(conn net.Conn) {
		input := bufio.NewScanner(conn)
		for input.Scan() {
			scanning <- input.Text()
		}
		// NOTE: ignoring potential errors from input.Err()
	}(conn)
	timeout := time.After(time.Second * timeoutsec)

	// Scanもゴルーチン化して、timeoutしない限りは新たなタイムアウト時間をセットし続ける
	for isTimeout := false; isTimeout == false; {
		select {
		case str := <-scanning:
			timeout = time.After(time.Second * timeoutsec)
			messages <- who + ": " + str
		case <-timeout:
			isTimeout = true
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

//!-main
