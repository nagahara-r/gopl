// Copyright © 2017 Yuki Nagahara
// 練習13-4: os.Exec を使用したbzip圧縮関数を作成します。

package bzip

import (
	"bytes"
	"io"
	"os/exec"
)

type writer struct {
	w io.Writer // underlying output stream
}

// NewWriter returns a writer for bzip2-compressed streams.
func NewWriter(out io.Writer) io.Writer {
	w := &writer{w: out}
	return w
}

func (w *writer) Write(data []byte) (int, error) {
	cmd := exec.Command("bzip2")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return 0, err
	}
	written, err := io.Copy(stdin, bytes.NewReader(data))
	if err != nil {
		return int(written), err
	}
	stdin.Close()
	out, err := cmd.Output()
	if err != nil {
		return int(written), err
	}

	for writesize := 0; writesize < len(out); {
		n, err := w.w.Write(out[writesize:])
		if err != nil {
			return n, err
		}
		writesize += n
	}

	return int(written), nil
}
