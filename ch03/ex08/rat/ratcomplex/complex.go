package ratcomplex

import (
	"math"
	"math/big"
)

// RatComplex はRatで複素数を実現します。
type RatComplex struct {
	real big.Rat
	imag big.Rat
}

// NewRatComplex はRat複素数を作成します。
func NewRatComplex(real float64, imag float64) (result RatComplex) {
	result.real.SetFloat64(real)
	result.imag.SetFloat64(imag)

	return result
}

// Add はRat複素数を足し算します。
func Add(a RatComplex, b RatComplex) (result RatComplex) {
	result.real.Add(&a.real, &b.real)
	result.imag.Add(&a.imag, &b.imag)
	return result
}

// Mul はRat複素数を掛け算します。
func Mul(a RatComplex, b RatComplex) (result RatComplex) {
	// 実数部
	result.real.Mul(&a.real, &b.real)
	result.imag.Mul(&a.imag, &b.imag)
	result.real.Sub(&result.real, &result.imag)

	// 虚数部
	imag1 := result.imag.Mul(&a.imag, &b.real)
	imag2 := result.imag.Mul(&a.real, &b.imag)
	result.imag.Add(imag1, imag2)

	return result
}

//new Complex(re * c.re - im * c.im, im * c.re + re * c.im)

// Abs はRat複素数の絶対値を取得します。
func Abs(z RatComplex) (result float64) {
	a, _ := z.real.Float64()
	b, _ := z.imag.Float64()
	return math.Hypot(a, b)
}
