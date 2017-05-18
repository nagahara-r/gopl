package main

import (
	"fmt"

	"github.com/naga718/golang-practice/ch04/ex07/reverse"
)

func main() {
	s := []byte("あaいbうcえdお")

	fmt.Println("reverse前: ", string(s))

	reverse.Reverse(s)
	fmt.Println("reverse後: ", string(s))
}
