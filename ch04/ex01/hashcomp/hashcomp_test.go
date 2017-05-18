package hashcomp

import (
	"crypto/sha256"
	"testing"
)

func TestSHA256XORCount(t *testing.T) {
	want := 125

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	result := SHA256XORCount(c1, c2)

	if result != want {
		t.Errorf("want = %v, result = %v", want, result)
	}
}

func TestSHA256XORCount2(t *testing.T) {
	want := 1
	c1 := [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3}
	c2 := [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}

	result := SHA256XORCount(c1, c2)

	if result != want {
		t.Errorf("want = %v, result = %v", want, result)
	}
}
