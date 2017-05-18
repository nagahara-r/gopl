package main

import (
	"fmt"
	"math/big"

	"github.com/naga718/golang-practice/ch03/ex13/bytesize"
)

func main() {
	fmt.Printf("KB = %v\n", bytesize.KB)
	fmt.Printf("MB = %v\n", bytesize.MB)
	fmt.Printf("GB = %v\n", bytesize.GB)
	fmt.Printf("TB = %v\n", bytesize.TB)
	fmt.Printf("PB = %v\n", bytesize.PB)
	fmt.Printf("EB = %v\n", bytesize.EB)
	fmt.Printf("ZB = %v\n", getZB())
	fmt.Printf("YB = %v\n", getYB())
}

func getZB() *big.Int {
	a := big.NewInt(bytesize.ZB / bytesize.MB)
	b := big.NewInt(bytesize.MB)

	x := big.NewInt(0)
	x.Mul(a, b)
	return x
}

func getYB() *big.Int {
	a := big.NewInt(bytesize.YB / bytesize.MB)
	b := big.NewInt(bytesize.MB)

	x := big.NewInt(0)
	x.Mul(a, b)

	return x
}
