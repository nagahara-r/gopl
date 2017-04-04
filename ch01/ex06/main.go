package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// 虹
var palette = []color.Color{
	color.White,
	color.RGBA{0xFF, 0x00, 0x00, 0xFF},
	color.RGBA{0xFF, 0xA5, 0x00, 0xFF},
	color.RGBA{0xFF, 0xFF, 0x00, 0xFF},
	color.RGBA{0x00, 0x80, 0x00, 0xFF},
	color.RGBA{0x00, 0xFF, 0xFF, 0xFF},
	color.RGBA{0x00, 0x00, 0xFF, 0xFF},
	color.RGBA{0x80, 0x00, 0x80, 0xFF},
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 発振器xが完了する周回数
		res     = 0.001 // 回転の分解能
		size    = 100   // 画像キャンバスは [-size..+size] の範囲を扱う
		nframes = 64    // アニメーションフレーム数
		delay   = 8     // 10ms 単位でのフレーム間の遅延

		loopn  = 35000 // 1枚の描写に使用するループ回数
		colorn = 7     // カラー数
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		j := 0
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			j++
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8((j/(loopn/colorn))+1))
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: エンコード無視
}
