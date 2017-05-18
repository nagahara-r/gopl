package main

import (
	"fmt"

	"github.com/naga718/golang-practice/ch03/ex12/anagram"
)

func main() {
	an1 := "TEST"
	an2 := "TTSE"

	fmt.Printf("Anagram is %v\n", anagram.IsAnagram(an1, an2))
}
