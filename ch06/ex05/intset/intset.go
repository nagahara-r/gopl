// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習6-5 uint64 を使わず intset を実装します。

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

const (
	uintSize = 32 << (^uint(0) >> 63)
	// 32bit なら 0bit shiftで 32 のまま。 64bitなら1bit shiftして64になる。
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/uintSize, uint(x%uintSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/uintSize, uint(x%uintSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith は2つのセットの積集合を計算します。
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

// DifferenceWith は s と i が一致しない値だけを取得します。
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			// twordのビット反転かつ論理積を取る：swordにあってtwordにない数だけが残る
			s.words[i] &= ^tword
		}
	}
}

// SymmetricDifference は2つのセットの排他集合を計算します。
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", uintSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

// Len は要素数を返します。
func (s *IntSet) Len() (count int) {
	for _, word := range s.words {
		for ; word > 0; count++ {
			word = word & (word - 1)
		}
	}
	return
}

// Remove はセットから x を取り除きます。
func (s *IntSet) Remove(x int) {
	if x < 0 {
		return
	}

	word, bit := x/uintSize, uint(x%uintSize)
	if word > len(s.words)-1 {
		return
	}
	s.words[word] &= ^(1 << bit)
}

// Clear はセットからすべての要素を取り除きます。
func (s *IntSet) Clear() {
	s.words = []uint{}
}

// Copy はセットのコピーを返します。
func (s *IntSet) Copy() *IntSet {
	newIntSet := IntSet{}
	newIntSet.words = make([]uint, len(s.words))
	copy(newIntSet.words, s.words)
	return &newIntSet
}

// AddAll は可変長引数にセットされた値をすべてリストに追加します。
func (s *IntSet) AddAll(x ...int) {
	for _, v := range x {
		s.Add(v)
	}
}

// Elems は rangeループで取得するに優れた []int を返します。
func (s *IntSet) Elems() (elems []int) {
	// var buf bytes.Buffer
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, 64*i+j)
				// fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	return
}
