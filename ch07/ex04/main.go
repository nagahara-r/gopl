// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習7-4：HTMLパーサが文字列からの入力を受け取れるio.Readerラッパーの作成。

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/net/html"
)

// StringReader は、String用のio.Readerです。
type StringReader string

func (sr *StringReader) Read(p []byte) (n int, err error) {
	// 全部コピーし終えているのであれば、EOFを返す
	if *sr == "" {
		return 0, io.EOF
	}

	// 全部コピー
	copy(p, []byte(*sr))
	*sr = ""
	return len(p), nil
}

// NewReader は StringReader を生成します。
func NewReader(str string) io.Reader {
	sr := StringReader(str)
	return &sr
}

func main() {
	// 便宜上、byteをstringに変換します。
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	str := string(b)

	doc, err := html.Parse(NewReader(str))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
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
