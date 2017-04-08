package langeconv

import (
	"testing"
)

func TestFtTom(t *testing.T) {
	want := Metre(0.3048)
	get := FtToM(1)

	if get != want {
		t.Errorf("FtTom = %v, want = %v", get, want)
	}
}

func TestMToFt(t *testing.T) {
	want := Feet(3.28)
	get := MToFt(1)

	if get != want {
		t.Errorf("MToFt = %v, want = %v", get, want)
	}
}
