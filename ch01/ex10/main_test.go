package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Args = []string{"", "http://gopl.io", "http://google.com"}
	main()
}
