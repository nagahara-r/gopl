package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

//func TestMain(t *testing.T) {
//	main()
//	statusCode := fetch("http://localhost:8000/")

//	if statusCode != 200 {
//		t.Fail()
//	}
//}

func TestLissajous(t *testing.T) {
	lissajous(os.Stdout, 20)
}

func fetch(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("StatusCode: ", resp.Status)
	_, err = io.Copy(os.Stdout, resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	return resp.StatusCode
}
