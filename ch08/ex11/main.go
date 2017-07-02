// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習8-11：複数のサーバに同時に問い合わせを行い、一つのリクエストが到着後すぐさまキャンセルする
// fetchの実装

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type httpGetTask struct {
	req      *http.Request
	resp     *http.Response
	cancelCh chan struct{}
}

func main() {
	fetch(os.Args[1:]...)
}

func fetch(urls ...string) {
	resp, err := getInParallel(urls...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", resp.Request.URL, err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
}

func getInParallel(urls ...string) (resp *http.Response, err error) {
	var httpGetTasks []httpGetTask
	respCh := make(chan *httpGetTask)

	for _, url := range urls {
		var req *http.Request
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			return
		}
		cancelCh := make(chan struct{})
		req.Cancel = cancelCh
		task := httpGetTask{req, nil, cancelCh}
		httpGetTasks = append(httpGetTasks, task)

		go get(&task, respCh)
	}

	for range urls {
		task := <-respCh
		if task != nil {
			resp = task.resp
			closeParallelRequests(task, httpGetTasks)
			break
		}
	}
	if resp == nil {
		err = fmt.Errorf("all request failed")
	}

	return
}

// closeParallelRequests は、最速でレスポンスしたリクエスト以外をCloseします
func closeParallelRequests(getTask *httpGetTask, httpGetTasks []httpGetTask) {
	for _, task := range httpGetTasks {
		if getTask.req.URL != task.req.URL {
			log.Printf("close() = %v", task.req.URL)
			close(task.cancelCh)
		}
	}
}

func get(task *httpGetTask, respCh chan *httpGetTask) {
	resp, err := http.DefaultClient.Do(task.req)
	if err != nil {
		log.Printf("get() error = %v", err)
		respCh <- nil
		return
	}
	task.resp = resp
	respCh <- task
}

//!-
