package main

import "testing"

func TestComma1(t *testing.T) {
	need := "1,234,567"

	if commaWithBuffer("1234567") != need {
		t.Errorf("need = %v, comma() = %v", need, commaWithBuffer("1234567"))
	}

}

func TestComma2(t *testing.T) {
	need := "1"

	if commaWithBuffer("1") != need {
		t.Errorf("need = %v, comma() = %v", need, commaWithBuffer("1"))
	}

}
