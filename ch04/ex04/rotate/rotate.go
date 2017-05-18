// Copyright © 2017 Yuki Nagahara
// rotate は指定の回数だけ左にintスライスを回転させます。

package rotate

func Rotate(a []int, r int) {
	r = r % len(a)

	buf := make([]int, len(a[:r]))
	copy(buf, a[:r])
	copy(a, a[r:len(a)])
	copy(a[len(a)-r:], buf)
}
