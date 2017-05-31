// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習5-13 ページの複製を行うようcrawlを修正

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	"gopl.io/ch5/links"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	for _, u := range list {
		if isSameDomain(u, url) {
			//log.Printf("same: %v", u)
			b, err := fetch(u)
			if err != nil {
				log.Printf("%v", err)
				continue
			}

			err = fileWrite(u, b)
			if err != nil {
				log.Printf("%v", err)
				continue
			}
		}
	}

	return list
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

//!-crawl

//!+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	//breadthFirst(crawl, os.Args[1:])

	for _, url := range os.Args[1:] {
		crawl(url)
	}
}

//!-main
