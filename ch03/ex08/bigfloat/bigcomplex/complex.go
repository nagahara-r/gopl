package bigcomplex

import (
	"math"
	"math/big"
)

// BigFloatComplex はBigFloatで複素数を実現します。
type BigFloatComplex struct {
	Real big.Float
	Imag big.Float
}

// NewBigFloatComplex はBigFloatComplexを作成します。
func NewBigFloatComplex(real float64, imag float64) (result BigFloatComplex) {
	result.Real.SetFloat64(real)
	result.Imag.SetFloat64(imag)

	return result
}

// Add はBigFloatComplexを足し算します。
func (BigFloatComplex) Add(a BigFloatComplex, b BigFloatComplex) (result BigFloatComplex) {
	result.Real.Add(&a.Real, &b.Real)
	result.Imag.Add(&a.Imag, &b.Imag)
	return result
}

// Mul はBigFloatComplexを掛け算します。
func (BigFloatComplex) Mul(a BigFloatComplex, b BigFloatComplex) (result BigFloatComplex) {
	// 実数部
	result.Real.Mul(&a.Real, &b.Real)
	result.Imag.Mul(&a.Imag, &b.Imag)
	result.Real.Sub(&result.Real, &result.Imag)

	// 虚数部
	imag := BigFloatComplex{}
	imag.Real.Mul(&a.Imag, &b.Real)
	imag.Imag.Mul(&a.Real, &b.Imag)
	imag.Imag.Add(&imag.Real, &imag.Imag)
	result.Imag = imag.Imag

	return result
}

//new Complex(re * c.re - im * c.im, im * c.re + re * c.im)

// Abs はBigFloatComplexの絶対値を取得します。
func (BigFloatComplex) Abs(z BigFloatComplex) (result float64) {
	a, _ := z.Real.Float64()
	b, _ := z.Imag.Float64()
	return math.Hypot(a, b)
}
