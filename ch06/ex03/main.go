// Copyright © 2017 Yuki Nagahara

package main

import (
	"fmt"

	"github.com/naga718/golang-practice/ch06/ex03/intset"
)

func main() {
	iset := intset.IntSet{}
	iset.AddAll(0, 1)

	target := intset.IntSet{}
	target.AddAll(1, 2, 3)

	fmt.Printf("intset = %v\n", iset.String())
	fmt.Printf("target = %v\n", target.String())

	iset.IntersectWith(&target)
	fmt.Printf("IntersectWith(&target) = %v\n", iset.String())

	iset.AddAll(0, 1)
	iset.DifferenceWith(&target)
	fmt.Printf("DifferenceWith(&target) = %v\n", iset.String())

	iset.AddAll(0, 1)
	iset.SymmetricDifference(&target)
	fmt.Printf("SymmetricDifference(&target) = %v\n", iset.String())
}
