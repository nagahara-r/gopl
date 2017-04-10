// Dup2 は入力に2回以上現れた行の数とその行のテキストを表示します。
// 標準入力から読み込むか、名前が指定されたファイルの一覧から読み込みます。
package main

import (
	"bufio"
	"fmt"
	"os"
)

type DupValue struct {
	Files []string
	Count int
}

func main() {
	counts := make(map[string]DupValue)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.Count > 1 {
			fmt.Printf("%d\t%s\tfiles=%v\n", n.Count, line, n.Files)
		}
	}
}

func countLines(f *os.File, counts map[string]DupValue) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		dupValue := counts[input.Text()]
		dupValue.Count++
		if contains(dupValue.Files, f.Name()) == false {
			dupValue.Files = append(dupValue.Files, f.Name())
		}
		counts[input.Text()] = dupValue
	}
	// 注意: input.Err() からのエラーの可能性を無視している
}

func contains(strslice []string, str string) bool {
	for _, v := range strslice {
		if str == v {
			return true
		}
	}
	return false
}
