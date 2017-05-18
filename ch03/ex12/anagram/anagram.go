// anagram は２つの文字列がアナグラムになっているか検査します。
// Copyright © 2017 Yuki Nagahara

package anagram

import "log"

// IsAnagram は２つの文字列がアナグラムになっているか検査します。
func IsAnagram(a string, b string) bool {
	return isAnagramRunes([]rune(a), []rune(b))
}

func isAnagramRunes(a []rune, b []rune) bool {
	if a == nil || b == nil {
		// 公開関数ではないのでNULLは通常来ない
		log.Fatal("Rune Slice nil Detected!")
	}

	// 同じ長さではない＝アナグラムではない
	if len(a) != len(b) {
		return false
	}

	if len(a) == 1 && len(b) == 1 {
		// 最後の一文字が一致しているか？
		return a[0] == b[0]
	}

	r := a[0]

	for i := range b {
		if r == b[i] {
			a = a[1:]
			b = append(b[:i], b[i+1:]...)
			return isAnagramRunes(a, b)
		}
	}

	// ここまで来るということは一致文字列がなかったとみなす。
	return false
}
