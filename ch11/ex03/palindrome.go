// Copyright © 2017 Yuki Nagahara

package palindrome

// IsPalindrome は列sが回文であるかを報告します。
func IsPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}

	return true
}
