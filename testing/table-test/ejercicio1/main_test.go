package main

import "testing"

func TestSuma(t *testing.T) {
	type prueba struct {
		datos     []int
		respuesta int
	}
	pruebas := []prueba{
		{[]int{1, 4, 6}, 11},
		{[]int{4, 3, 7}, 14},
		{[]int{1, 8, 6}, 15},
		{[]int{1, 0, 1}, 2},
	}
	for _, p := range pruebas {
		v := suma(p.datos...)
		if v != p.respuesta {
			t.Error("Expected", p.respuesta, "Got", v)
		}
	}

}
