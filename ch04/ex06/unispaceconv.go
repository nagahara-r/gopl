// Copyright © 2017 Yuki Nagahara
// unispaceconv は[]byteスライス内で隣接するUnicodespaceをASCIIスペースに圧縮します。

package unispaceconv

import (
	"unicode"
	"unicode/utf8"
)

// ConvertUnicodeSpaceToASCII は[]byteスライス内で隣接するUnicodespaceをASCIIスペースに圧縮します。
func ConvertUnicodeSpaceToASCII(strs []byte) []byte {
	compsize := 0

	for i := range strs {
		if !utf8.RuneStart(strs[i]) {
			continue
		}

		r, size := utf8.DecodeRune(strs[i:])
		if unicode.IsSpace(r) {
			strs[i] = ' '
			copy(strs[i+1:], strs[i+size:])
			compsize += size - 1
		}
	}

	return strs[:len(strs)-compsize]
}
