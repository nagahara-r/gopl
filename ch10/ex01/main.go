// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習10-1: jpeg変換プログラムを拡張して、入力と出力をサポートしているすべての画像形式で行えるようにします。

// See page 287.

// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"strings"
)

var (
	encoder = map[string]func(image.Image, io.Writer) error{
		"jpeg": toJPEG,
		"jpg":  toJPEG,
		"png":  toPNG,
		"gif":  toGIF,
	}
)

func main() {
	f := flag.String("f", "jpg", "format : (jpeg, jpg, png, gif)")
	flag.Parse()
	*f = strings.ToLower(*f)

	img, kind, err := image.Decode(os.Stdin)
	if err != nil {
		log.Fatalf("decode(): %v\n", err)
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)

	enc, ok := encoder[*f]
	if !ok {
		log.Fatalf("encoder(): invalid format: %v\n", *f)
	}

	if err := enc(img, os.Stdout); err != nil {
		log.Fatalf("encode(): %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(img image.Image, out io.Writer) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toPNG(img image.Image, out io.Writer) error {
	return png.Encode(out, img)
}

func toGIF(img image.Image, out io.Writer) error {
	return gif.Encode(out, img, &gif.Options{})
}
