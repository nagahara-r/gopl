// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習8-6: 平行なクローラに深さ制限を追加します。

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

type seenLinks struct {
	url   string
	depth int
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!+
func main() {
	maxDepth := flag.Int("depth", -1, "Depth Limit (minus value presents unlimited)")
	flag.Parse()

	worklist := make(chan []seenLinks)  // lists of URLs, may have duplicates
	unseenLinks := make(chan seenLinks) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- storeLinks(os.Args[2:], -1) }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				if link.depth > *maxDepth && *maxDepth > -1 {
					continue
				}
				foundLinks := crawl(link.url)
				depth := link.depth
				go func() { worklist <- storeLinks(foundLinks, depth) }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.url] {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
}

func storeLinks(linkStrs []string, currentDepth int) (sLinks []seenLinks) {
	for _, link := range linkStrs {
		sLinks = append(sLinks, seenLinks{link, currentDepth + 1})
	}
	return
}

//!-
