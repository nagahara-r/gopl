// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習7-3：ツリー内の値を見せるStringメソッドの作成

// See page 101.

// Package treesort provides insertion sort using an unbalanced binary tree.
package treesort

import (
	"strconv"
	"strings"
)

//!+
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
// * main関数から確認できるようにtreeを返却するよう変更
func Sort(values []int) *tree {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	return root
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() (str string) {
	return strings.Join([]string{"[", buildStr(t), "]"}, "")
}

func buildStr(t *tree) (str string) {
	if t != nil {
		str = strings.Join([]string{str, buildStr(t.left)}, " ")
		str = strings.Join([]string{str, strconv.Itoa(t.value)}, " ")
		str = strings.Join([]string{str, buildStr(t.right)}, " ")
	}
	return strings.TrimSpace(str)
}

//!-
