package rotate

import "testing"

func TestRotate(t *testing.T) {
	want := [...]int{2, 3, 4, 5, 0, 1}
	a := [...]int{0, 1, 2, 3, 4, 5}

	rotate(a[:], 2)

	if want != a {
		t.Errorf("want = %v, rotate = %v", want, a)
	}
}
