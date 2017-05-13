package reverse

import (
	"testing"
	"unsafe"
)

func TestReverse(t *testing.T) {
	want := [...]int{5, 4, 3, 2, 1, 0}
	a := [...]int{0, 1, 2, 3, 4, 5}

	// ポインタになると要素数がわからないので、要素数を教えてあげる
	reverse((*[]int)(unsafe.Pointer(&a)), len(a))

	if want != a {
		t.Errorf("want = %v, reverse = %v", want, a)
	}

}
