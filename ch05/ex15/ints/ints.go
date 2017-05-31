package ints

// Copyright © 2017 Yuki Nagahara
// 練習5-15: 可変長引数のMax Minの実装

// Max は与えられた引数のなかで最も大きな数を返します。
// 引数が 0 個のときは、intの初期値である0を返します。
func Max(vals ...int) (max int) {
	for i, val := range vals {
		if i == 0 {
			max = val
		}

		if max < val {
			max = val
		}
	}

	return
}

// Min は与えられた引数のなかで最も小さな数を返します。
// 引数が 0 個のときは、intの初期値である0を返します。
func Min(vals ...int) (min int) {
	for i, val := range vals {
		if i == 0 {
			min = val
		}

		if min > val {
			min = val
		}
	}

	return
}
