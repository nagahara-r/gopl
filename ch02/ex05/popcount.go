package popcount

// pc[i] は i のポピュレーションカウント
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount はxのポピュレーションカウント（1が設定されているビット数）を返します。
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountByLoop(x uint64) int {
	popCount := 0
	var i uint64
	for i = 0; i < 8; i++ {
		popCount += int(pc[byte(x>>(i*8))])
	}

	return popCount
}

func PopCountByBitshift(x uint64) int {
	popCount := 0
	var i uint64
	for i = 0; i < 64; i++ {
		popCount += int((x >> i) & 0x1)
	}

	return popCount
}

func PopCountByBitclear(x uint64) int {
	var i int
	for i = 0; x > 0; i++ {
		x = x & (x - 1)
	}

	return i
}
