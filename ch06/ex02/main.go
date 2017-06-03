package main

import (
	"fmt"

	"./intset"
)

func main() {
	iset := intset.IntSet{}
	iset.AddAll(0, 1, 2, 3, 128)

	fmt.Printf("intset = %v\n", iset.String())
	iset.Remove(128)
	fmt.Printf("Remove(128) = %v\n", iset.String())
}
