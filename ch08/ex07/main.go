// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習8-7: Webサイトのローカルなミラーを平行に作成します。

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
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"sync"
	"time"

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
				go func() {
					worklist <- storeLinks(foundLinks, depth)
				}()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)

	timeout := time.After(time.Second * 10)

	dling := new(sync.WaitGroup)

	for {
		select {
		case list := <-worklist:
			for _, link := range list {
				timeout = time.After(time.Second * 10)
				if !seen[link.url] {
					seen[link.url] = true
					unseenLinks <- link

					dling.Add(1)
					go func(link string) {
						fileDownload(link, os.Args[2:])
						dling.Done()
					}(link.url)
				}
			}
		case <-timeout:
			// どのチャネルも10秒応答しなくなったらダウンロードしているゴルーチンを待って終了
			fmt.Println("Timed out")
			dling.Wait()
			return
		}
	}
}

func fileDownload(link string, doms []string) {
	for _, dom := range doms {
		if isSameDomain(link, dom) {
			log.Printf("download: %v", link)
			b, err := fetch(link)
			if err != nil {
				log.Printf("err: %v", err)
				continue
			}

			fileWrite(link, b)
		}
	}
}

func fetch(url string) (b []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return
	}
	b, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		return
	}
	//fmt.Printf("%s", b)
	return
}

func fileWrite(u string, b []byte) (err error) {
	usrc, err := url.Parse(u)
	if err != nil {
		return
	}

	fpath := usrc.Host + usrc.Path
	_, fn := path.Split(usrc.Path)
	if fn == "" {
		fpath = fpath + "index.html"
	}

	file, err := createFileWithMkdir(fpath)
	defer file.Close()
	if err != nil {
		return
	}

	file.Write(b)
	return
}

// ファイルを作成します。ディレクトリがない場合はあわせて作ります。
func createFileWithMkdir(p string) (file *os.File, err error) {
	log.Printf("create: %v", p)
	dir, _ := path.Split(p)
	_, err = os.Stat(dir)

	if err != nil {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return
		}
	}

	return os.Create(p)
}

// 同じドメインかどうかを調べます。
func isSameDomain(dst string, src string) bool {
	usrc, err := url.Parse(src)
	if err != nil {
		log.Printf("(src) %v", err)
		return false
	}

	udst, err := url.Parse(dst)
	if err != nil {
		log.Printf("(dst) %v", err)
		return false
	}

	return udst.Host == usrc.Host
}

func storeLinks(linkStrs []string, currentDepth int) (sLinks []seenLinks) {
	for _, link := range linkStrs {
		sLinks = append(sLinks, seenLinks{link, currentDepth + 1})
	}
	return
}

//!-
