package main

import (
	"fmt"

	"github.com/naga718/golang-practice/ch04/ex03/reverse"
)

func main() {
	a := [...]int{5, 4, 3, 2, 1, 0}

	fmt.Printf("reverse前: %v\n", a)

	// ポインタになると要素数がわからないので、要素数を教えてあげる
	reverse.Reverse(&a[0], len(a))

	fmt.Printf("reverse後: %v\n", a)

}
