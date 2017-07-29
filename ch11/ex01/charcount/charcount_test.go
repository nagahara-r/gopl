// Copyright © 2017 Yuki Nagahara
// 練習11-1: Charcount のテスト

package charcount

import (
	"io"
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCharCount(t *testing.T) {
	tests := []struct {
		input   io.Reader
		counts  map[rune]int
		utflen  [utf8.UTFMax + 1]int
		invalid int
	}{
		{
			strings.NewReader("abcde"),
			map[rune]int{rune('a'): 1, rune('b'): 1, rune('c'): 1, rune('d'): 1, rune('e'): 1},
			[utf8.UTFMax + 1]int{0, 5},
			0,
		}, {
			strings.NewReader("あいうえおaa"),
			map[rune]int{rune('a'): 2, rune('あ'): 1, rune('い'): 1, rune('う'): 1, rune('え'): 1, rune('お'): 1},
			[utf8.UTFMax + 1]int{0, 2, 0, 5},
			0,
		}, {
			strings.NewReader("あa\n "),
			map[rune]int{rune('あ'): 1, rune('a'): 1, rune('\n'): 1, rune(' '): 1},
			[utf8.UTFMax + 1]int{0, 3, 0, 1},
			0,
		}, {
			strings.NewReader("あa\n \xff\xff"),
			map[rune]int{rune('あ'): 1, rune('a'): 1, rune('\n'): 1, rune(' '): 1},
			[utf8.UTFMax + 1]int{0, 3, 0, 1},
			2,
		},
		{
			strings.NewReader(""),
			map[rune]int{},
			[utf8.UTFMax + 1]int{},
			0,
		},
		{
			nil,
			nil,
			[utf8.UTFMax + 1]int{},
			0,
		},
	}

	for _, test := range tests {
		counts, utflen, invalid := CharCount(test.input)
		if !reflect.DeepEqual(counts, test.counts) {
			t.Errorf("counts: \n%v\nexcepted: \n%v\n", counts, test.counts)
		}

		if !reflect.DeepEqual(utflen, test.utflen) {
			t.Errorf("counts: \n%v\nexcepted: \n%v\n", utflen, test.utflen)
		}

		if invalid != test.invalid {
			t.Errorf("invalid: \n%v\nexcepted: \n%v\n", invalid, test.invalid)
		}
	}
}
