// Fetchall はURLを平行に取り出し、時間と大きさを表示します
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for i, url := range os.Args[1:] {
		go fetch(url, ch, i) // ゴルーチンを開始
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // ch チャネルから受信
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, num int) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // ch チャネルへ送信
		return
	}

	file, err := os.Create("./" + strconv.Itoa(num) + ".html")
	if err != nil {
		ch <- fmt.Sprintf("while fileopen: %v", err)
		return
	}

	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close() // 資源をリークさせない
	file.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
