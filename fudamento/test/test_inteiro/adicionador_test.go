package test_inteiro

import "testing"

func TestAdicionador(t *testing.T){

	soma := Adicionado(2, 2)
	esperado := 4

	if soma != esperado {
		t.Errorf("esperado '%d', resultado '%d'", esperado, soma)
	}
	
}