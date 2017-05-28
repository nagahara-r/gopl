// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習5-11: トポロジーソート 循環報告版

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"log"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

//!+main
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(key string)

	visitAll = func(key string) {
		if m[key] == nil {
			return
		}

		circkey, det := searchCirculation(key, m)
		if det {
			log.Fatalf("Circulation Reported!\n%v: %v\n%v: %v\n", key, m[key], circkey, m[circkey])
		}

		for _, v := range m[key] {
			if !seen[v] {
				seen[v] = true
				visitAll(v)
				order = append(order, v)
			}
		}

		if !seen[key] {
			seen[key] = true
			order = append(order, key)
		}
	}

	for key := range m {
		visitAll(key)
	}

	return order
}

// serachCirculation は、トポロジーソート中に循環がないか探して報告します。
func searchCirculation(key string, m map[string][]string) (string, bool) {
	if m[key] == nil {
		return "", false
	}

	for _, v := range m[key] {
		if m[v] == nil {
			continue
		}

		for _, vv := range m[v] {
			if key == vv {
				return v, true
			}
		}
	}

	return "", false
}

//!-main
