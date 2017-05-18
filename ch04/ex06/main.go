package main

import (
	"fmt"

	"github.com/naga718/golang-practice/ch04/ex06/unispaceconv"
)

func main() {
	a := []byte("全角　スペース")

	fmt.Printf("全角スペース置き換え前: %v\n", string(a))

	a = unispaceconv.ConvertUnicodeSpaceToASCII(a)

	fmt.Printf("全角スペース置き換え後: %v\n", string(a))
}
