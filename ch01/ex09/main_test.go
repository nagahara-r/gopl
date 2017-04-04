package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Args = []string{"", "http://gopl.io"}
	main()
}

func TestMainNoScheme(t *testing.T) {
	os.Args = []string{"", "gopl.io"}
	main()
}
