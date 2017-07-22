package main

// Copyright © 2017 Yuki Nagahara
// 練習9-4: 任意の数のゴルーチンをチャネルで接続します。

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	r := flag.Int64("r", 2, "Number of Goroutine")
	flag.Parse()
	if *r <= 1 {
		log.Fatal("Number of Goroutine must be 2 or more")
	}

	//go printMemPeriodically(1)

	pipes := make([]chan struct{}, *r)
	for i := range pipes {
		pipes[i] = make(chan struct{})
	}

	var prev chan struct{}
	for _, pipe := range pipes {
		if prev == nil {
			prev = pipe
			continue
		}
		go pipeline(prev, pipe)
		prev = pipe
	}

	// Start Send Pipe
	log.Println("Start Send Pipe")
	ts := time.Now()
	pipes[0] <- struct{}{}
	<-pipes[len(pipes)-1]

	fmt.Printf("Processing Time: %v\n", time.Now().Sub(ts))
}

func pipeline(receive chan struct{}, send chan struct{}) {
	send <- <-receive
}

func printMemPeriodically(second int64) {
	tick := time.Tick(time.Second * time.Duration(second))
	var mem runtime.MemStats
	for {
		select {
		case <-tick:
			runtime.ReadMemStats(&mem)
			log.Printf("alloc = %v, Total alloc = %v, Heap alloc = %v, Heap sys = %v\n", mem.Alloc, mem.TotalAlloc, mem.HeapAlloc, mem.HeapSys)
		}
	}
}
