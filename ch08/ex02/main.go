// Copyright © 2017 Yuki Nagahara
// 練習8-2: 並行に動作するFTPサーバを作成します。

package main

import (
	"flag"
	"log"
	"net"
	"strconv"

	"github.com/naga718/golang-practice/ch08/ex02/ftpd"
)

func main() {
	port := flag.Int("port", 8000, "port number")
	u := flag.String("u", "anonymous", "user")
	p := flag.String("p", "", "password")
	flag.Parse()

	user := ftpd.User{*u, *p}

	listener, err := net.Listen("tcp4", ":"+strconv.Itoa(*port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}

		go ftpd.HandleConn(conn, user) // handle connections concurrently
	}
	//!-
}
