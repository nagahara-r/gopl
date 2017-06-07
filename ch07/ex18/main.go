// Copyright © 2017 Yuki Nagahara
// 練習7-18 XMLドキュメントのノードツリーを構築するプログラム

package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("input: \n%v\n", string(b))

	dec := xml.NewDecoder(bytes.NewReader(b))
	nodes := parseNode(dec, nil)

	printNodes(nodes)
}

func printNodes(nodes []Node) {
	for _, node := range nodes {
		switch n := node.(type) {
		case CharData:
			fmt.Printf("%v\n", n)
		case Element:
			fmt.Printf("Element: %v\n", n.Type.Local)
			for _, at := range n.Attr {
				fmt.Printf("Attr: key = %v, value = %v\n", at.Name.Local, at.Value)
			}
			printNodes(n.Children)
		}
	}
}

func parseNode(dec *xml.Decoder, nodes []Node) []Node {
	tok, err := dec.Token()
	if err == io.EOF {
		return nodes
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
		os.Exit(1)
	}
	switch tok := tok.(type) {
	case xml.StartElement:
		el := Element{}

		el.Type = tok.Name
		el.Attr = tok.Attr
		el.Children = parseNode(dec, nil)
		nodes = append(nodes, el)
		break
	case xml.EndElement:
		return nodes
	case xml.CharData:
		cd := CharData(tok)
		nodes = append(nodes, cd)
		break
	}

	nodes = parseNode(dec, nodes)

	return nodes
}
