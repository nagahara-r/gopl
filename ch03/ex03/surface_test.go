package main

import (
	"strings"
	"testing"
)

func TestSurface(t *testing.T) {
	s := getSurface()

	if strings.Contains(s, "NaN") {
		t.Error("NaN value detected")
	}
}
