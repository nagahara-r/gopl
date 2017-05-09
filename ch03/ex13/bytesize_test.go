package bytesize

import (
	"math/big"
	"testing"
)

// キロバイト
func TestKBsize(t *testing.T) {
	//fmt.Printf("%v\n", KB)

	if KB != 8000 {
		t.Errorf("want = 8000, get = %v", KB)
	}
}

// メガバイト
func TestMBsize(t *testing.T) {
	if MB != 8000000 {
		t.Errorf("want = 8000000, get = %v", MB)
	}
}

// ギガバイト
func TestGBsize(t *testing.T) {
	if GB != 8000000000 {
		t.Errorf("want = 8000000000, get = %v", GB)
	}
}

// テラバイト
func TestTBsize(t *testing.T) {
	if TB != 8000000000000 {
		t.Errorf("want = 8000000000000, get = %v", TB)
	}
}

// ペタバイト
func TestPBsize(t *testing.T) {
	if PB != 8000000000000000 {
		t.Errorf("want = 8000000000000000, get = %v", PB)
	}
}

// エクサバイト
func TestEBsize(t *testing.T) {
	a := big.NewInt(EB / MB)
	b := big.NewInt(MB)

	x := big.NewInt(0)
	x.Mul(a, b)
	//fmt.Printf("%v\n", x.String())

	if x.String() != "8000000000000000000" {
		t.Errorf("want = 8000000000000000000, get = %v", x.String())
	}
}

// ゼタバイト
func TestZBsize(t *testing.T) {
	a := big.NewInt(ZB / MB)
	b := big.NewInt(MB)

	x := big.NewInt(0)
	x.Mul(a, b)
	//fmt.Printf("%v\n", x.String())

	if x.String() != "8000000000000000000000" {
		t.Errorf("want = 8000000000000000000000, get = %v", x.String())
	}
}

// ヤタバイト
func TestYBsize(t *testing.T) {
	a := big.NewInt(YB / MB)
	b := big.NewInt(MB)

	x := big.NewInt(0)
	x.Mul(a, b)
	//fmt.Printf("%v\n", x.String())

	if x.String() != "8000000000000000000000000" {
		t.Errorf("want = 8000000000000000000000000, get = %v", x.String())
	}
}
