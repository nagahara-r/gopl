package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/naga718/golang-practice/ch04/ex01/hashcomp"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("x と X のSHA256で異なるビット数は %v\n", hashcomp.SHA256XORCount(c1, c2))
}
