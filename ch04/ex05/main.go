package main

import (
	"fmt"

	"github.com/naga718/golang-practice/ch04/ex05/delstring"
)

func main() {
	a := []string{"apple", "orange", "orange", "grape", "pineapple", "grape", "grape"}

	fmt.Printf("連続重複削除前: %v\n", a)

	a = delstring.DeleteSideDuplicate(a)

	fmt.Printf("連続重複削除後: %v\n", a)
}
