// Copyright © 2017 Yuki Nagahara
// 練習7-2 バイト数を数えるWriterの実装

package main

import (
	"bytes"
	"fmt"
	"io"
)

// ByteCountWriter はバイト数を保持する構造体です。
type ByteCountWriter struct {
	writer io.Writer
	size   int64
}

// CountingWriter は書き込んだバイト数を保持するWriterを返します。
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	bcw := ByteCountWriter{}
	bcw.writer = w
	return &bcw, &bcw.size
}

func (bcw *ByteCountWriter) Write(p []byte) (int, error) {
	size, err := bcw.writer.Write(p)
	bcw.size += int64(size)
	return size, err
}

func main() {
	w := new(bytes.Buffer)
	bcw, size := CountingWriter(w)

	bcw.Write([]byte("hello"))
	fmt.Printf("words = %v, size = %v\n", string(w.Bytes()), *size)

	bcw.Write([]byte(" world"))
	fmt.Printf("words = %v, size = %v\n", string(w.Bytes()), *size)
}
