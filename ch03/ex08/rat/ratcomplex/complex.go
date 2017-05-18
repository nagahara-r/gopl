package ratcomplex

import (
	"math"
	"math/big"
)

// RatComplex はRatで複素数を実現します。
type RatComplex struct {
	Real big.Rat
	Imag big.Rat
}

// NewRatComplex はRat複素数を作成します。
func NewRatComplex(real float64, imag float64) (result RatComplex) {
	result.Real.SetFloat64(real)
	result.Imag.SetFloat64(imag)

	return result
}

// Add はRat複素数を足し算します。
func Add(a RatComplex, b RatComplex) (result RatComplex) {
	result.Real.Add(&a.Real, &b.Real)
	result.Imag.Add(&a.Imag, &b.Imag)
	return result
}

// Mul はRat複素数を掛け算します。
func Mul(a RatComplex, b RatComplex) (result RatComplex) {
	// 実数部
	result.Real.Mul(&a.Real, &b.Real)
	result.Imag.Mul(&a.Imag, &b.Imag)
	result.Real.Sub(&result.Real, &result.Imag)

	// 虚数部
	imag := RatComplex{}
	imag.Real.Mul(&a.Imag, &b.Real)
	imag.Imag.Mul(&a.Real, &b.Imag)
	imag.Imag.Add(&imag.Real, &imag.Imag)
	result.Imag = imag.Imag

	return result
}

//new Complex(re * c.re - im * c.im, im * c.re + re * c.im)

// Abs はRat複素数の絶対値を取得します。
func Abs(z RatComplex) (result float64) {
	a, _ := z.Real.Float64()
	b, _ := z.Imag.Float64()
	return math.Hypot(a, b)
}
