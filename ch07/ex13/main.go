package main

import (
	"fmt"
	"log"

	"./eval"
)

func main() {
	exprs := []string{
		"sqrt(A / pi)",
		"pow(x, 3) + pow(y, 3)",
		"pow(x, 3) + pow(y, 3)",
		"5 / 9 * (F - 32)",
		"-1 + -x",
		"-1 - x",
	}

	for _, e := range exprs {
		ex, err := eval.Parse(e)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(e + " =>")
		fmt.Println(ex.String())
	}
}
