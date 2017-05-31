// Copyright © 2017 Yuki Nagahara
// 練習7.10: 列が回文であるかを報告する関数 IsPalindrome の実装

package palindrome

import "sort"

// IsPalindrome は列sが回文であるかを報告します。
func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}

	return true
}
