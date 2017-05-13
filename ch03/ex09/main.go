// Mandelbrot emits a PNG image of the Mandelbrot fractal.
// 課題3-9: マンデルブロ集合をPNGイメージで表示させるWebサーバ
// Usage: http://localhost:8000/?x=0&y=0&scale=1
// http://localhost:8000/?x=0&y=256&scale=10000000
package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	listenAddress := "localhost:8000"

	handler := func(w http.ResponseWriter, r *http.Request) {
		xq := fixNullTo0(r.URL.Query().Get("x"))
		yq := fixNullTo0(r.URL.Query().Get("y"))
		scaleq := fixNullTo1(r.URL.Query().Get("scale"))

		x, err := strconv.Atoi(xq)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "400 bad request\n error: Invalid x value")
			return
		}

		y, err := strconv.Atoi(yq)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "400 bad request\n error: Invalid y value")
			return
		}

		scale, err := strconv.ParseFloat(scaleq, 64)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "400 bad request\n error: Invalid scale value")
			return
		}

		result, err := fractal(x, y, scale)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "400 bad request\n error = %v", err.Error())
		} else {
			w.Write(result)
		}
	}
	http.HandleFunc("/", handler)

	fmt.Println("Start Server: ", listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, nil))

}

func fixNullTo0(str string) string {
	if str == "" {
		return "0"
	}
	return str
}

func fixNullTo1(str string) string {
	if str == "" {
		return "1"
	}
	return str
}

func fractal(xfix int, yfix int, scale float64) (writer []byte, err error) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := (float64(py)+float64(yfix)*scale)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := (float64(px)+float64(xfix)*scale)/width*(xmax-xmin) + xmin
			z := complex(x, y) / complex(scale, 0)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}

	pngbuffer := new(bytes.Buffer)
	err = png.Encode(pngbuffer, img)

	return pngbuffer.Bytes(), err
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette.Plan9[n]
		}
	}
	return color.Black
}
