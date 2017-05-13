// Copyright © 2017 Yuki Nagahara
// 標準入力で得た文章内の単語の出現頻度を報告します。

package main

import (
	"bufio"
	"fmt"
	"os"
)

type wordCounter struct {
	word  string
	count int
}

// count はwordの単語が出た回数をカウントします。
func count(word string, wordList []wordCounter) (result []wordCounter) {
	e, ok := search(word, wordList)

	if ok {
		e.count++
	} else {
		w := wordCounter{word, 1}
		wordList = append(wordList, w)
	}

	return wordList
}

// serach はすでに登録されたwordがあるか探します。
func search(word string, wordList []wordCounter) (entity *wordCounter, ok bool) {
	for i := range wordList {
		if word == wordList[i].word {
			return &wordList[i], true
		}
	}

	return nil, false
}

// printWordCount は数えた単語数を表示します。
func printWordCount(wordList []wordCounter) {
	fmt.Printf("word\tcount\n")
	for _, w := range wordList {
		fmt.Printf("%v\t%v\n", w.word, w.count)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	wordList := []wordCounter{}

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		w := scanner.Text()
		wordList = count(w, wordList)
	}

	printWordCount(wordList)
}
