// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習5-17: ElementsByTagNameの実装

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	nodes := ElementByTagName(doc, "img", "a")

	fmt.Printf("%v nodes detected\n\n", len(nodes))
	for _, node := range nodes {
		fmt.Printf("node = %v\n", node)
	}
	//!-call
	return nil
}

// ElementByTagName は指定されたノードから指定されたタグを含むノードを探します。
func ElementByTagName(doc *html.Node, name ...string) (nodes []*html.Node) {
	for _, tname := range name {
		nodes = append(nodes, forEachNode(doc, tname, startElement, endElement)...)
	}
	return
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, tname string, pre, post func(n *html.Node, id string) bool) (nodes []*html.Node) {
	if pre != nil {
		if pre(n, tname) {
			nodes = append(nodes, n)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, forEachNode(c, tname, pre, post)...)
	}

	// preElementが見つかれば要素確定なので、post分は不要
	// if post != nil {
	// 	if post(n, tname) {
	// 		nodes = append(nodes, n)
	// 	}
	// }

	return nodes
}

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node, tname string) bool {
	if n.Type == html.ElementNode {
		// log.Printf("%v", n.Data)
		if n.Data == tname {
			return true
		}
	}
	return false
}

func endElement(n *html.Node, tname string) bool {
	if n.Type == html.ElementNode {
		if n.Data == tname {
			return true
		}
	}
	return false
}

//!-startend
