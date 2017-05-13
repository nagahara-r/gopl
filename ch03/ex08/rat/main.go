// Mandelbrot emits a PNG image of the Mandelbrot fractal.
// ratで計算します。
package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"os"

	"github.com/naga718/golang-practice/ch03/ex08/rat/ratcomplex"
)

func main() {
	ratFractal()
}

func ratFractal() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		xfix                   = 0
		yfix                   = 256
		scale                  = 10000000
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := (float64(py)+yfix*scale)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := (float64(px)+xfix*scale)/width*(xmax-xmin) + xmin
			scalerat := ratcomplex.NewRatComplex(float64(1)/float64(scale), float64(0))
			z := ratcomplex.Mul(ratcomplex.NewRatComplex(x, y), scalerat)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z ratcomplex.RatComplex) color.Color {
	const iterations = 3
	const contrast = 15

	var v ratcomplex.RatComplex
	for n := uint8(0); n < iterations; n++ {
		//log.Println("n = ", n)
		v = ratcomplex.Mul(v, v)
		v = ratcomplex.Add(v, z)
		if ratcomplex.Abs(v) > 2 {
			return palette.Plan9[n]
		}
	}
	return color.Black
}
