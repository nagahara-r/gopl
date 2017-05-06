// commaWithBuffer は負ではない10進表記整数文字列のカンマを挿入します。
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("-1234567.89"))
}

func comma(s string) string {
	str, isNegative := trimNegativeSymbol(s)

	strs := strings.Split(str, ".")

	var buf bytes.Buffer
	n := 0

	for i := 0; i <= len(strs[0]); i = i + 3 {
		if i == 0 {
			n = (len(strs[0]) - i) % 3
			fmt.Fprintf(&buf, "%v", strs[0][i:i+n])
			i = n
		} else {
			if buf.Len() > 0 {
				buf.WriteByte(',')
			}
			fmt.Fprintf(&buf, "%v", strs[0][n:i])
			n = n + 3
		}

	}

	str = buf.String()

	if len(strs) > 1 {
		str = str + "." + strs[1]
	}

	if isNegative {
		str = "-" + str
	}

	return str
}

func trimNegativeSymbol(s string) (string, bool) {
	if s[0] == '-' {
		return s[1:], true
	} else {
		return s, false
	}
}
