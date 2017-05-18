package main

import (
	"fmt"

	"github.com/naga718/golang-practice/ch04/ex04/rotate"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}

	fmt.Printf("回転前: %v\n", a)

	rotate.Rotate(a[:], 2)

	fmt.Printf("左2回転後: %v\n", a)
}
