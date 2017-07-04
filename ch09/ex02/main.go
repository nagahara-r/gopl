package main

import (
	"fmt"

	"github.com/naga718/golang-practice/ch09/ex02/popcount"
)

func main() {
	fmt.Printf("PopCount(0) = %v\n", popcount.PopCount(0))
	fmt.Printf("PopCount(1) = %v\n", popcount.PopCount(1))
	fmt.Printf("PopCount(9223372036854775807) = %v\n", popcount.PopCount(9223372036854775807))
}
