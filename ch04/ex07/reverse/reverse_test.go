package reverse

import "testing"

func TestReverseUTF8(t *testing.T) {
	want := "おえういあ"
	get := []byte("あいうえお")
	Reverse(get)

	if string(get) != want {
		t.Errorf("want = %v, get = %v", want, string(get))
	}
}

func TestReverseAscii(t *testing.T) {
	want := "edcba"
	get := []byte("abcde")
	Reverse(get)

	if string(get) != want {
		t.Errorf("want = %v, get = %v", want, string(get))
	}
}

func TestReverseAsciiAndUTF8(t *testing.T) {
	want := "おdえcうbいaあ"
	get := []byte("あaいbうcえdお")
	Reverse(get)

	if string(get) != want {
		t.Errorf("want = %v, get = %v", want, string(get))
	}
}
