// reverse は配列を逆転します。
// Copyright © 2017 Yuki Nagahara

package reverse

import "unsafe"

const (
	memAlign = 8
)

func reverse(s *[]int, l int) {
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		*getPointer(s, i), *getPointer(s, j) = *getPointer(s, j), *getPointer(s, i)
	}

}

// getPointer は*[]intで指定された配列の要素をintポインタにして返します。
func getPointer(addr *[]int, suffix int) (result *int) {
	// unsafe.Pointer にして、アドレスを足し算可能なように uintptr にキャスト
	uip := uintptr(unsafe.Pointer(addr))

	// 1. インデックス分の移動 -> index = suffix * memAlign
	// 2. 配列のアドレスへ足し算 address = uip + index
	// 3. unsafe.Pointer化 unsp = unsafe.Pointer(address)
	// 4. int アドレス型にキャスト return (*int)(unsp)
	return (*int)(unsafe.Pointer(uip + uintptr(suffix*memAlign)))
}
