package main

import (
	"fmt"

	"./intset"
)

func main() {
	iset := intset.IntSet{}
	iset.Add(0)
	iset.Add(1)
	iset.Add(2)
	iset.Add(3)

	fmt.Printf("intset = %v\n", iset.String())
	fmt.Printf("Len() = %v\n", iset.Len())
	iset.Remove(1)
	fmt.Printf("Remove(1) = %v\n", iset.String())
	cpis := iset.Copy()
	fmt.Printf("Copy() = %v\n", cpis.String())
	iset.Clear()
	fmt.Printf("Clear() = %v\n", iset.String())

}
