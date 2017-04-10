package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Args = []string{"main.go", "test1.txt", "test2.txt", "test3.txt"}
	main()
}

func TestMainWithFileError(t *testing.T) {
	os.Args = []string{"main.go", "test1.txt", "err.txt"}
	main()
}
