package main

// Copyright © 2017 Yuki Nagahara

import (
	"fmt"

	"./ints"
)

func main() {
	a := []int{-12345, -1234, 0, 123, 12345}

	fmt.Printf("array = %v\n", a)
	fmt.Printf("max = %v\n", ints.Max(a...))
	fmt.Printf("min = %v\n", ints.Min(a...))
}
