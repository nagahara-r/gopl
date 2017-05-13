package main

import "testing"

func TestReverseUTF8(t *testing.T) {
	want := "おえういあ"
	get := []byte("あいうえお")
	reverse(get)

	if string(get) != want {
		t.Errorf("want = %v, get = %v", want, string(get))
	}
}

func TestReverseAscii(t *testing.T) {
	want := "edcba"
	get := []byte("abcde")
	reverse(get)

	if string(get) != want {
		t.Errorf("want = %v, get = %v", want, string(get))
	}
}

func TestReverseAsciiAndUTF8(t *testing.T) {
	want := "おdえcうbいaあ"
	get := []byte("あaいbうcえdお")
	reverse(get)

	if string(get) != want {
		t.Errorf("want = %v, get = %v", want, string(get))
	}
}
