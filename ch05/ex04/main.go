// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習5.4 visit を拡張し、通常リンク、画像、スクリプト、スタイルシート　のリンクを取得します。

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

//!-main

// visitはリンク、外部スタイル、JavaScript、画像のリンクをHTMLノードから探し
// 一覧表示します。
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "link") {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	} else if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "img") {
		for _, script := range n.Attr {
			if script.Key == "src" {
				links = append(links, script.Val)
			}
		}
	}

	// ノード下降可能か
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}

	// ノードに次の要素があるか
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	return links
}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
