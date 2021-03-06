// newton
// ニュートンフラクタルを表示させます。
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			// 4つの根に近似する値ごとに色付けとグラデーションする
			if (real(z) - 1) > -0.5 {
				return color.RGBA{255, 0, 0, 255 - contrast*i}
			} else if (real(z) + 1) < 0.5 {
				return color.RGBA{0, 255, 0, 255 - contrast*i}
			} else if (imag(z) - 1) > -0.5 {
				return color.RGBA{0, 0, 255, 255 - contrast*i}
			} else {
				return color.RGBA{255, 255, 0, 255 - contrast*i}
			}
		}
	}
	return color.Black
}
