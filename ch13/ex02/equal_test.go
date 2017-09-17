// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習13-2: 循環を検査する関数を作成します。

package cycle

import "testing"

func TestEqual(t *testing.T) {
	var uncyclepointer bool
	type CyclePtr *CyclePtr
	var cyclePtr1, cyclePtr2 CyclePtr
	cyclePtr1 = &cyclePtr2
	cyclePtr2 = &cyclePtr1

	type CycleSlice []CycleSlice
	var cycleSlice = make(CycleSlice, 1)
	cycleSlice[0] = cycleSlice

	for _, test := range []struct {
		i    interface{}
		want bool
	}{
		// Uncycle slice
		{[]string{"a", "b", "c"}, false},
		// Uncycle Pointer
		{&uncyclepointer, false},
		// slice cycle
		{cycleSlice, true},
		// pointer cycle
		{cyclePtr1, true},
		// Invalid
		{nil, false},
	} {
		if Circulate(test.i) != test.want {
			t.Errorf("Circulate(%v) = %v, but expect %v",
				test.i, !test.want, test.want)
		}
	}
}
