package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)
	if total != 30 {
		t.Errorf("Soma incorreta, resultado esperado %d, resultado %d.", 30, total)
	}
}
