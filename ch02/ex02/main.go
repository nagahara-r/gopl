package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/naga718/golang-practice/ch02/ex01"
	"github.com/naga718/golang-practice/ch02/ex02/langeconv"
	"github.com/naga718/golang-practice/ch02/ex02/weightconv"
)

func main() {
	var args []string

	if len(os.Args) <= 1 {
		var arg string
		fmt.Scan(&arg)
		args = []string{"", arg}
	} else {
		args = os.Args
	}

	for _, arg := range args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)

		m := langeconv.Metre(t)
		ft := langeconv.Feet(t)

		lb := weightconv.Pound(t)
		kg := weightconv.Kilo(t)

		fmt.Printf("%s = %s, %s = %s \n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Printf("%s = %s, %s = %s \n", m, langeconv.MToFt(m), ft, langeconv.FtToM(ft))
		fmt.Printf("%s = %s, %s = %s \n", lb, weightconv.LbToKg(lb), kg, weightconv.KgToLb(kg))
	}
}
