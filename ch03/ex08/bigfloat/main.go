// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/naga718/golang-practice/ch03/ex08/bigfloat/bigcomplex"
)

func main() {
	bigMandelbrot()
}

func bigMandelbrot() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		xfix                   = 0
		yfix                   = 256
		scale                  = 10000000
	)

	start := time.Now()

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := (float64(py)+yfix*scale)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := (float64(px)+xfix*scale)/width*(xmax-xmin) + xmin
			scalebigfloat := bigcomplex.NewBigFloatComplex(float64(1)/float64(scale), float64(0))
			z := scalebigfloat.Mul(bigcomplex.NewBigFloatComplex(x, y), scalebigfloat)

			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}

	end := time.Now()
	log.Printf("big.Float = %f Seconds\n", (end.Sub(start)).Seconds())

	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z bigcomplex.BigFloatComplex) color.Color {
	const iterations = 200
	const contrast = 15

	var v bigcomplex.BigFloatComplex
	for n := uint8(0); n < iterations; n++ {
		//log.Println("n = ", n)
		v = v.Mul(v, v)
		v = v.Add(v, z)
		if v.Abs(v) > 2 {
			// ループ回数が多い＝赤
			// ループ回数が少ない＝青
			// 中間＝緑
			return color.RGBA{contrast * n, contrast * 2 * n, 255 - contrast*n, 255}
		}
	}
	return color.Black
}
