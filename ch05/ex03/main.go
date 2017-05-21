// Copyright © 2017 Yuki Nagahara
// HTMLドキュメント内のすべてのテキストノードの内容を表示します。
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("outline: %v\n", err)
	}
	viewtextnode(doc)
}

func viewtextnode(n *html.Node) {
	if n.Type == html.TextNode {
		if !isInvisibleScript(n.Parent) && !(strings.TrimSpace(n.Data) == "") {
			fmt.Printf("%v\n", n.Data)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		viewtextnode(c)
	}
}

// テキストがScriptノード内かどうか調べる
func isInvisibleScript(n *html.Node) bool {
	if n == nil {
		return false
	}

	if strings.ToLower(n.Data) == "script" || strings.ToLower(n.Data) == "style" {
		//log.Printf("Script!: %v", n.Data)
		return true
	}

	return false
}
