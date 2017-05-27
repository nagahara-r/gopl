package main

import (
	"log"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestForEachNode(t *testing.T) {
	var tests = []struct {
		input    string
		id       string
		expected string
	}{
		{
			"<html><head><title>Hoge</title></head><body><br><p></body></html>",
			"test",
			"",
		},
		{
			"<html><head><title>Hoge</title></head><body><br><div id='test'>div</a><p></body></html>",
			"test",
			"div",
		},
		{
			"<!-- This is Comment --><html><head><title>Hoge<aaaa></title></head><body><br><div id=''>Link</a><p></body></html>",
			"",
			"div",
		},
		{
			"<script type='javascript'>this is java script</script><html><head><title>Hoge<aaaa></title></head><body><br><div id=''><a href='link.html'>Link</a><p></body></html>",
			"test",
			"",
		},
	}

	for i, test := range tests {
		n, err := html.Parse(strings.NewReader(test.input))
		if err != nil {
			t.Errorf("%v", err)
		}

		e := ElementbyID(n, test.id)
		data := ""

		log.Printf("%v", e)

		if e != nil {
			data = e.Data
		}

		if data != test.expected {
			t.Errorf("test = %v, data = %v, expected = %v", i, data, test.expected)
		}
	}

}
