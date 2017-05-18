// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"time"
)

func main() {
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
			z := complex(x, y) / scale

			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}

	end := time.Now()
	log.Printf("complex128 = %f Seconds\n", (end.Sub(start)).Seconds())

	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			// ループ回数が多い＝赤
			// ループ回数が少ない＝青
			// 中間＝緑
			return color.RGBA{contrast * n, contrast * 2 * n, 255 - contrast*n, 255}
		}
	}
	return color.Black
}
