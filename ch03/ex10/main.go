// commaWithBuffer は負ではない10進表記整数文字列のカンマを挿入します。
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(commaWithBuffer("123456789"))
}

func commaWithBuffer(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer

	for i := 0; i <= len(s); i = i + 3 {
		if i == 0 {
			n = (len(s) - i) % 3
			fmt.Fprintf(&buf, "%v", s[i:i+n])
			i = n
		} else {
			if buf.Len() > 0 {
				buf.WriteByte(',')
			}
			fmt.Fprintf(&buf, "%v", s[n:i])
			n = n + 3
		}
	}

	return buf.String()
}
