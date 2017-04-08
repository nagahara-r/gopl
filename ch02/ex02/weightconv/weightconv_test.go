package weightconv

import (
	"testing"
)

func TestKgToLb(t *testing.T) {
	want := Pound(2.2)
	get := KgToLb(1)

	if get != want {
		t.Errorf("KgToLb = %v, want = %v", get, want)
	}
}

func TestLbToKg(t *testing.T) {
	want := Kilo(0.45)
	get := LbToKg(1)

	if get != want {
		t.Errorf("LbToKg = %v, want = %v", get, want)
	}
}
