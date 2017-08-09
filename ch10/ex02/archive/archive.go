// Copyright © 2017 Yuki Nagahara
// 練習10-2: ZIPとTARファイルを扱える汎用のアーカイブ読み込み関数を定義します。

package archive

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Unzip はファイルを解凍します。
// file には解凍するファイルを指定し、dst には解凍先のディレクトリを指定します。
// 戻り値は 解凍に成功したディレクトリの struct File です。
func Unzip(zipfile string, dir string) (err error) {
	f, ok := sniff(zipfile)
	if !ok {
		return fmt.Errorf("invalid format")
	}

	err = f.unzip(zipfile, dir)
	return
}

// 下記は image/format.go を参考に作成しました。
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// RegisterFormat はFormatを登録します。
func RegisterFormat(name, magic string, unzip func(string, string) error) {
	formats = append(formats, format{name, magic, unzip})
}

// A format holds an image format's name, magic header and how to decode it.
type format struct {
	name, magic string
	unzip       func(string, string) error
}

var formats []format

// Sniff determines the format of r's data.
func sniff(filepath string) (format, bool) {
	for _, f := range formats {
		file, err := os.Open(filepath)
		if err != nil {
			return format{}, false
		}
		lr := io.LimitReader(file, int64(len(f.magic)))
		b, err := ioutil.ReadAll(lr)
		if err == nil && match(f.magic, b) {
			file.Close()
			return f, true
		}
		err = file.Close()
		if err != nil {
			return format{}, false
		}
	}
	return format{}, false
}

// Match reports whether magic matches b. Magic may contain "?" wildcards.
func match(magic string, b []byte) bool {
	if len(magic) != len(b) {
		return false
	}
	for i, c := range b {
		if magic[i] != c && magic[i] != '?' {
			return false
		}
	}
	return true
}
