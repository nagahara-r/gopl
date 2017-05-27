// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習5-8: ElementByID はノード中のid一致を見つけ、そのノードを返します。

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"log"
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
	node := ElementbyID(doc, "lowframe")
	log.Printf("node = %v", node)
	//!-call

	return nil
}

// ElementbyID は指定されたノードから指定されたIDを含むノードを探します。
func ElementbyID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, startElement, endElement)
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	if pre != nil {
		if pre(n, id) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		det := forEachNode(c, id, pre, post)
		if det != nil {
			return det
		}
	}

	if post != nil {
		if post(n, id) {
			return n
		}
	}

	return nil
}

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node, id string) bool {

	if n.Type == html.ElementNode {
		for _, at := range n.Attr {
			if at.Key == "id" && at.Val == id {
				return true
			}
		}
	}

	return false
}

func endElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, at := range n.Attr {
			if at.Key == "id" && at.Val == id {
				return true
			}
		}
	}

	return false
}

//!-startend
