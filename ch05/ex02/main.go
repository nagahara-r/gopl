// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// HTMLドキュメントツリー内でその要素名と要素の数を対応させるマッピングを行います。
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

//!+
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	nmap := nodemap(nil, doc)

	fmt.Printf("%v\n", nmap)
}

func nodemap(nmap map[string]int, n *html.Node) map[string]int {
	if nmap == nil {
		nmap = map[string]int{}
	}

	if n.Type == html.ElementNode {
		nmap[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodemap(nmap, c)
	}

	return nmap
}

//!-
