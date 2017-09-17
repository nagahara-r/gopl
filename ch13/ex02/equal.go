// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習13-2: 循環を検査する関数を作成します。

package cycle

import (
	"reflect"
	"unsafe"
)

func circulate(v reflect.Value, seen map[vtype]bool) bool {
	if !v.IsValid() {
		return false
	}

	// cycle check
	if v.CanAddr() {
		vptr := unsafe.Pointer(v.UnsafeAddr())
		vt := vtype{vptr, v.Type()}
		if seen[vt] {
			return true // already seen
		}
		seen[vt] = true
	}

	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return circulate(v.Elem(), seen)

	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			return circulate(v.Index(i), seen)
		}
	}

	// とくに循環を検出しなかった場合
	return false
}

//!-

// Circulate は循環を検査します。
func Circulate(i interface{}) bool {
	seen := make(map[vtype]bool)
	return circulate(reflect.ValueOf(i), seen)
}

type vtype struct {
	v unsafe.Pointer
	t reflect.Type
}

//!-comparison
