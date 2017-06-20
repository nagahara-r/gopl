// Copyright © 2017 Yuki Nagahara

// 練習8-1: ClockWall は複数の時計サーバにアクセスし、並行して時間を表示します。
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

// ClockRegion には時計の表示リージョン名、サーバアドレスを指定します
type ClockRegion struct {
	Region string
	Dest   string
}

var (
	m = new(sync.Mutex)
)

func main() {
	crs := parse()
	if len(crs) < 1 {
		log.Fatalf("接続対象がありません。")
	}

	// 画面クリア
	fmt.Fprint(os.Stderr, "\033[2J")
	defer fmt.Fprint(os.Stderr, "\033[2J")

	for i, cr := range crs {
		go startClock(i, cr.Region, cr.Dest)
	}

	// EOF、またはInterruptまで待つ
	ioutil.ReadAll(os.Stdin)
}

func parse() (crs []ClockRegion) {
	for _, arg := range os.Args {
		ss := strings.Split(arg, "=")
		if len(ss) != 2 {
			log.Printf("parse error = %v\n", arg)
			continue
		}
		crs = append(crs, ClockRegion{ss[0], ss[1]})
	}

	return
}

func startClock(index int, region string, dest string) {
	conn, err := net.Dial("tcp", dest)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		str, err := readTime(conn)
		if err != nil {
			log.Fatalf("%v", err)
		}
		printTime(region, str, (index*3)+1)
	}
}

func printTime(region string, str string, printIndex int) {
	m.Lock()
	defer m.Unlock()
	fmt.Fprintf(os.Stderr, "\033[%v;0H", printIndex)
	fmt.Fprintf(os.Stderr, "%v\n%v", region, str)
}

func readTime(conn net.Conn) (str string, err error) {
	buf := make([]byte, 128)

	for n := 0; n == 0; {
		n, err = conn.Read(buf)
	}

	if err != nil {
		return // e.g., client disconnected
	}

	str = string(buf)
	return
}
