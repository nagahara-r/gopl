package main

import (
	"bytes"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestForEachNode(t *testing.T) {
	var tests = []struct {
		input        string
		expectedData []string
	}{
		{
			"<html><head><title>Hoge</title></head><body><br><p></body></html>",
			[]string{"html", "head", "title", "Hoge", "body", "br", "p"},
		},
		{
			"<html><head><title>Hoge</title></head><body><br><a href='link.html'>Link</a><p></body></html>",
			[]string{"html", "head", "title", "Hoge", "body", "br", "a", "Link", "p"},
		},
		{
			"<!-- This is Comment --><html><head><title>Hoge<aaaa></title></head><body><br><a href='link.html'>Link</a><p></body></html>",
			[]string{"This is Comment", "html", "head", "title", "Hoge<aaaa>", "body", "br", "a", "Link", "p"},
		},
		{
			"<script type='javascript'>this is java script</script><html><head><title>Hoge<aaaa></title></head><body><br><a href='link.html'>Link</a><p></body></html>",
			[]string{"html", "head", "script", "this is java script", "title", "Hoge<aaaa>", "body", "br", "a", "Link", "p"},
		},
		{
			"",
			[]string{"html", "head", "body"},
		},
	}

	for _, test := range tests {
		n, err := html.Parse(strings.NewReader(test.input))
		if err != nil {
			t.Errorf("%v", err)
		}

		w := &bytes.Buffer{}

		forEachNode(w, n, startElement, endElement)

		n, err = html.Parse(bytes.NewReader(w.Bytes()))
		if err != nil {
			t.Errorf("%v", err)
		}

		_, datas := parseHTML(n, nil, nil)

		//
		if !sliceEquals(datas, test.expectedData) {
			t.Errorf("forEachNode = \n%v, expected = \n%v", datas, test.expectedData)
		}
	}
}

func sliceEquals(slice []string, src []string) bool {
	if len(slice) != len(src) {
		return false
	}

	for i := range slice {
		if slice[i] != src[i] {
			return false
		}
	}
	return true
}

func parseHTML(n *html.Node, types []html.NodeType, datas []string) ([]html.NodeType, []string) {
	if types == nil {
		types = []html.NodeType{}
	}

	if datas == nil {
		datas = []string{}
	}

	types = append(types, n.Type)
	if strings.TrimSpace(n.Data) != "" {
		datas = append(datas, strings.TrimSpace(n.Data))
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		types, datas = parseHTML(c, types, datas)
	}

	return types, datas
}

func trimAllSpaces(str string) string {
	return strings.Replace(str, " ", "", -1)
}
