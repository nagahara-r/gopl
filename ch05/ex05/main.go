// Copyright © 2017 Yuki Nagahara
// 練習5.5 HTML内に含まれる単語と画像の数を返すCountWordsAndImagesの実装

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "CountWordsAndImages: %v\n", err)
			continue
		}

		fmt.Printf("words: %v, images: %v\n", words, images)
	}
}

// CountWordsAndImages は単語数と画像数を数えます。
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	documents, images := visit([]string{}, 0, n)
	words = countWords(strings.Join(documents, " "))

	return
}

func visit(documents []string, images int, n *html.Node) ([]string, int) {
	if n.Type == html.TextNode {
		if !isInvisibleScript(n.Parent) && !(strings.TrimSpace(n.Data) == "") {
			documents = append(documents, n.Data)
		}
	} else if n.Type == html.ElementNode && (n.Data == "img") {
		images++
	}

	// ノード下降可能か
	if n.FirstChild != nil {
		documents, images = visit(documents, images, n.FirstChild)
	}

	// ノードに次の要素があるか
	if n.NextSibling != nil {
		documents, images = visit(documents, images, n.NextSibling)
	}

	return documents, images
}

// テキストがScriptノード内かどうか調べる
func isInvisibleScript(n *html.Node) bool {
	if n == nil {
		return false
	}

	if strings.ToLower(n.Data) == "script" || strings.ToLower(n.Data) == "style" {
		return true
	}

	return false
}

func countWords(str string) (count int) {
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		count++
		//log.Printf("%v: %v", count, scanner.Text())
	}

	return
}
