package main

import (
	"fmt"
	"math/rand"

	"github.com/naga718/golang-practice/ch07/ex03/treesort"
)

func main() {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	tree := treesort.Sort(data)

	fmt.Printf("%v\n", tree.String())
}
