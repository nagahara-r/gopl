// Copyright © 2017 Yuki Nagahara

package delstring

import (
	"log"
	"testing"
)

func TestDeleteSideDuplicate(t *testing.T) {
	want := []string{"apple", "orange", "grape", "pineapple", "grape"}
	a := []string{"apple", "orange", "orange", "grape", "pineapple", "grape", "grape"}

	a = DeleteSideDuplicate(a)

	if !isSameSlice(want, a) {
		t.Errorf("want = %v, DeleteSideDuplicate = %v", want, a)
	}
}

func isSameSlice(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			log.Printf("Difference! a = %v b = %v", a[i], b[i])
			return false
		}
	}
	return true
}
