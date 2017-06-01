// Copyright © 2017 Yuki Nagahara
// 練習5-16: strings.Joinの可変長引数バージョンの実装

package strings

import "strings"

// Join は strings.Join の可変長引数バージョンです。
func Join(sep string, a ...string) (ret string) {
	for _, str := range a {
		ret = ret + str + sep
	}

	return strings.TrimSuffix(ret, sep)
}
