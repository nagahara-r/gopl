package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(s []byte) {
	if utf8.RuneCount(s) < 1 {
		return
	}
	r, size := utf8.DecodeLastRune(s)

	byteRightShift(s, size)
	copy(s[:size], string(r))
	reverse(s[size:])
}

func byteRightShift(s []byte, size int) {
	length := len(s) - 1

	// copyでやるとうまくいかない
	// スライスの前からコピーするため？
	for i := 0; length-size-i >= 0; i++ {
		s[length-i] = s[length-i-size]
	}
}

func main() {
	s := []byte("あaいbうcえdお")

	reverse(s)
	fmt.Println(string(s))
}
