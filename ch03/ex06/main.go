// SuperSamplingはピクセル化の影響を薄めるプログラムです。
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	// ファイル読み込み
	inputFile, err := os.Open("mandelbrot_x2.0.png")

	if nil != err {
		log.Println(err)
		return
	}
	inputImage, _, err := image.Decode(inputFile)

	if nil != err {
		log.Println(err)
	}

	defer inputFile.Close()

	// ファイル出力
	outputFile, err := os.Create("mandelbrot_fixed.png")
	if nil != err {
		log.Println(err)
	}

	outputImage := superSampling(inputImage)  // 変換
	err = png.Encode(outputFile, outputImage) // エンコード

	if nil != err {
		fmt.Println(err)
	}

	defer outputFile.Close()
}

func superSampling(inputImage image.Image) image.Image {
	rect := inputImage.Bounds()
	width := rect.Size().X
	height := rect.Size().Y
	rgba := image.NewRGBA(rect)

	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			var setColor color.RGBA
			// 座標(x,y)のR, G, B, α の値を取得
			r1, g1, b1, a1 := inputImage.At(x-1, y-1).RGBA()
			r2, g2, b2, a2 := inputImage.At(x-1, y).RGBA()
			r3, g3, b3, a3 := inputImage.At(x, y-1).RGBA()
			r4, g4, b4, a4 := inputImage.At(x, y).RGBA()

			// 4つの座標からRGBA値を取ってきて平均を取る
			r := (r1 + r2 + r3 + r4) / 4
			g := (g1 + g2 + g3 + g4) / 4
			b := (b1 + b2 + b3 + b4) / 4
			a := (a1 + a2 + a3 + a4) / 4
			setColor.R = uint8(r)
			setColor.G = uint8(g)
			setColor.B = uint8(b)
			setColor.A = uint8(a)
			rgba.Set(x, y, setColor)
		}
	}

	return rgba.SubImage(rect)

}
