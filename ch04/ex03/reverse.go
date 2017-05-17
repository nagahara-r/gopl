// reverse は配列を逆転します。
// Copyright © 2017 Yuki Nagahara

package reverse

import "unsafe"

const (
	intSize = 4 << (^uint(0) >> 63)
	// 32bit なら 0bit shiftで 4 のまま。 64bitなら1bit shiftして8になる。

	// Goでintが32ビットか64ビットか調べる方法
	// http://qiita.com/ruiu/items/28c77ed483cec365fe84
)

func reverse(s *int, l int) {
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		*getPointer(s, i), *getPointer(s, j) = *getPointer(s, j), *getPointer(s, i)
	}

}

// getPointer は*intで指定された配列の要素をintポインタにして返します。
func getPointer(addr *int, index int) (result *int) {
	// unsafe.Pointer にして、アドレスを足し算可能なように uintptr にキャスト
	uip := uintptr(unsafe.Pointer(addr))

	// インデックス分の移動 -> indexAddress = index * intSize
	indexAddress := uintptr(index * intSize)

	// 1. 配列のアドレスへ足し算 address = uip + indexAddress
	// 2. unsafe.Pointer化（しないと次のキャストができない） unsp = unsafe.Pointer(address)
	// 3. int ポインタ型にキャスト return (*int)(unsp)
	return (*int)(unsafe.Pointer(uip + indexAddress))
}
