package main

import "testing"

func TestSuma(t *testing.T) {
	v := suma(2, 2)
	if v != 4 {
		t.Error("Expected", 4, "Got", v)
	}
}
