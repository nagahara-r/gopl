// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習8-5: 並列実行による実行時間差の計測

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"runtime"
	"sync"
	"time"

	"github.com/naga718/golang-practice/ch03/ex08/bigfloat/bigcomplex"
)

func main() {
	bigMandelbrot()
}

func bigMandelbrot() {
	p := flag.Int("p", 1, "Number of Processor(s)")
	flag.Parse()
	//fmt.Printf("CPU=%d\n", *p)
	runtime.GOMAXPROCS(*p)
	stime := time.Now()

	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		xfix                   = 0
		yfix                   = 256
		scale                  = 10000000
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	wg := new(sync.WaitGroup)

	for py := 0; py < height; py++ {
		y := (float64(py)+yfix*scale)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := (float64(px)+xfix*scale)/width*(xmax-xmin) + xmin
			scalebigfloat := bigcomplex.NewBigFloatComplex(float64(1)/float64(scale), float64(0))
			z := scalebigfloat.Mul(bigcomplex.NewBigFloatComplex(x, y), scalebigfloat)

			// マンデルブロ計算時に並列化
			wg.Add(1)
			go func(z bigcomplex.BigFloatComplex) {
				// Image point (px, py) represents complex value z.
				img.Set(px, py, mandelbrot(z))
				wg.Done()
			}(z)
		}

		wg.Wait()
	}

	proctime := time.Now().Sub(stime)
	fmt.Printf("proccesor = %v, time = %v\n", *p, proctime)

	//png.Encode(os.Stdout, img) // NOTE: ignoring errors
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
