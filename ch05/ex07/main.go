// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習5-7: 汎用HTMLプリティプリンタ。各種エレメントノード、テキストノード、コメントノードを出力します。

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

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
	forEachNode(os.Stdout, doc, startElement, endElement)
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(w io.Writer, n *html.Node, pre, post func(w io.Writer, n *html.Node)) {
	if pre != nil {
		pre(w, n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(w, c, pre, post)
	}

	if post != nil {
		post(w, n)
	}
}

//!-forEachNode

//!+startend
var depth int

func startElement(w io.Writer, n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Fprintf(w, "%*s", depth*2, "")
		fmt.Fprintf(w, "<%s", n.Data)
		for _, att := range n.Attr {
			fmt.Fprintf(w, " %v='%v'", att.Key, att.Val)
		}

		if hasChildNode(n) {
			fmt.Fprintf(w, ">\n")
			depth++
		} else {
			fmt.Fprintf(w, "/>\n")
		}
		//fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
	} else if n.Type == html.TextNode {
		if !(strings.TrimSpace(n.Data) == "") {
			fmt.Fprintf(w, "%v\n", n.Data)
		}
	} else if n.Type == html.CommentNode {
		fmt.Fprintf(w, "<!-- %v -->\n", n.Data)
	}
}

func endElement(w io.Writer, n *html.Node) {
	if !hasChildNode(n) {
		return
	}

	if n.Type == html.ElementNode {
		depth--
		fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Data)
	}
}

func hasChildNode(n *html.Node) bool {
	return n.LastChild != nil
}

//!-startend
