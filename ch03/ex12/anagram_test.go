// Copyright © 2017 Yuki Nagahara

package anagram

import "testing"

func TestIsAnagramASCII(t *testing.T) {
	an1 := "TEST"
	an2 := "TTSE"

	if !IsAnagram(an1, an2) {
		t.Errorf("Anagram Check Failed")
	}
}

func TestIsAnagramUTF8(t *testing.T) {
	an1 := "UTF8文字列で　テスト"
	an2 := "テストUTF8　文字列で"

	if !IsAnagram(an1, an2) {
		t.Errorf("Anagram Check Failed")
	}
}

func TestIsAnagramBLonger(t *testing.T) {
	an1 := "テスト"
	an2 := "ストテtesttest"

	if IsAnagram(an1, an2) {
		t.Errorf("Anagram Check Failed")
	}
}

func TestIsAnagramALonger(t *testing.T) {
	an1 := "B is shorter than A"
	an2 := "A is longer than B"

	if IsAnagram(an1, an2) {
		t.Errorf("Anagram Check Failed")
	}
}
