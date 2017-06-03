// Copyright © 2017 Yuki Nagahara

package main

import (
	"fmt"

	"./intset"
)

func main() {
	iset := intset.IntSet{}
	iset.AddAll(1, 2, 3)

	fmt.Printf("Elems() = %v\n", iset.Elems())
}
