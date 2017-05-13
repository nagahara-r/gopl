package bigcomplex

import (
	"math"
	"math/big"
)

// BigFloatComplex はBigFloatで複素数を実現します。
type BigFloatComplex struct {
	real big.Float
	imag big.Float
}

// NewBigFloatComplex はBigFloatComplexを作成します。
func NewBigFloatComplex(real float64, imag float64) (result BigFloatComplex) {
	result.real.SetFloat64(real)
	result.imag.SetFloat64(imag)

	return result
}

// Add はBigFloatComplexを足し算します。
func (BigFloatComplex) Add(a BigFloatComplex, b BigFloatComplex) (result BigFloatComplex) {
	result.real.Add(&a.real, &b.real)
	result.imag.Add(&a.imag, &b.imag)
	return result
}

// Mul はBigFloatComplexを掛け算します。
func (BigFloatComplex) Mul(a BigFloatComplex, b BigFloatComplex) (result BigFloatComplex) {
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

// Abs はBigFloatComplexの絶対値を取得します。
func (BigFloatComplex) Abs(z BigFloatComplex) (result float64) {
	a, _ := z.real.Float64()
	b, _ := z.imag.Float64()
	return math.Hypot(a, b)
}
