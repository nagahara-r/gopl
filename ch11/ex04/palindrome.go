// Copyright © 2017 Yuki Nagahara

package palindrome

import "unicode"

// IsPalindrome は列sが回文であるかを報告します。
func IsPalindrome(s string) bool {
	s = removeSkipLiteral(s)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}

	return true
}

func removeSkipLiteral(s string) string {
	runes := []rune(s)
	var nrunes []rune

	for _, r := range runes {
		if !unicode.IsSpace(r) && !unicode.IsPunct(r) {
			nrunes = append(nrunes, r)
		}
	}

	return string(nrunes)
}
