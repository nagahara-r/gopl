// Copyright © 2017 Yuki Nagahara

package hashcomp

// SHA256XORCount は２つのSHAハッシュの異なるビット数を数えます。
func SHA256XORCount(x [32]byte, y [32]byte) (result int) {
	for i := range x {
		// 排他論理和を取り、異なるビットを抽出し、popCountで数える
		xor := x[i] ^ y[i]
		result += popCount(xor)
	}

	return result
}

func popCount(x byte) int {
	var i int
	for i = 0; x > 0; i++ {
		x = x & (x - 1)
	}

	return i
}
