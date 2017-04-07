package tempconv

import (
	"testing"
)

func TestCtoF(t *testing.T) {
	// 30°C = 86°F
	var want Fahrenheit
	want = 86
	ctof := CToF(30)
	if ctof != want {
		t.Errorf("CtoF = %v, want = %v", ctof, want)
	}
}

func TestFtoC(t *testing.T) {
	// 30°C = 86°F
	want := Celsius(30)
	ftoc := FToC(86)
	if ftoc != want {
		t.Errorf("FtoC = %v, want = %v", ftoc, want)
	}
}

func TestKtoC(t *testing.T) {
	// 0K = -273.15°C
	want := Celsius(-273.15)
	ktoc := KToC(0)
	if ktoc != want {
		t.Errorf("KtoC = %v, want = %v", ktoc, want)
	}
}

func TestCtoK(t *testing.T) {
	// 273.15K = 0°C
	want := Kelvin(273.15)
	ctok := CToK(0)
	if ctok != want {
		t.Errorf("CtoK = %v, want = %v", ctok, want)
	}
}

func TestKtoF(t *testing.T) {
	// 303.15K = 30°C = 86°F
	want := Fahrenheit(86)
	ktof := KToF(303.15)
	if ktof != want {
		t.Errorf("KtoF = %v, want = %v", ktof, want)
	}
}

func TestFtoK(t *testing.T) {
	// 303.15K = 30°C = 86°F
	want := Kelvin(303.15)
	ftok := FToK(86)
	if ftok != want {
		t.Errorf("FtoK = %v, want = %v", ftok, want)
	}
}
