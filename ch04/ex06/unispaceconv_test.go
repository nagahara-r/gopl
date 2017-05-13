// Copyright © 2017 Yuki Nagahara

package unispaceconv

import (
	"log"
	"testing"
)

func TestConvertUnicodeSpaceToASCII(t *testing.T) {
	want := []byte("Unicode ですので  全角スペース   使ってます")
	a := []byte("Unicode　ですので　　全角スペース　　　使ってます")

	a = ConvertUnicodeSpaceToASCII(a)

	if !isSameSlice(a, want) {
		t.Errorf("want = \n%v, \nDeleteSideDuplicate = \n%v", want, a)
	}
}

func isSameSlice(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			log.Printf("Difference! a = %v b = %v", a[i], b[i])
			return false
		}
	}
	return true
}
