// Copyright © 2017 Yuki Nagahara

package main

import (
	"fmt"

	"./strings"
)

func main() {
	a := []string{"a", "b", "cdefg"}
	sep := " & "

	fmt.Printf("%v\n", strings.Join(sep, a[0], a[1], a[2]))
}
