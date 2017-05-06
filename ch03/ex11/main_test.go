package main

import "testing"

func TestComma1(t *testing.T) {
	need := "1,234,567"

	if comma("1234567") != need {
		t.Errorf("need = %v, comma() = %v", need, comma("1234567"))
	}

}

func TestComma2(t *testing.T) {
	need := "-1,234,567"

	if comma("-1234567") != need {
		t.Errorf("need = %v, comma() = %v", need, comma("-1234567"))
	}
}

func TestComma3(t *testing.T) {
	need := "1,234,567.89"

	if comma("1234567.89") != need {
		t.Errorf("need = %v, comma() = %v", need, comma("1234567.89"))
	}
}

func TestComma4(t *testing.T) {
	need := "-1,234,567.89"

	if comma("-1234567.89") != need {
		t.Errorf("need = %v, comma() = %v", need, comma("-1234567.89"))
	}
}
