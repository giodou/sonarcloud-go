package main

import "testing"

func TestSoma(t *testing.T){
	total := Soma(15,15)
	
	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}

func TestSubtracao(t *testing.T){
	total := Subtracao(15,5)
	
	if total != 10 {
		t.Errorf("Resultado da Subtracao é inválido: Resultado %d. Esperado: %d", total, 10)
	}
}

func TestDivisao(t *testing.T){
	total := Divisao(15,5)
	
	if total != 3 {
		t.Errorf("Resultado da Divisao é inválido: Resultado %d. Esperado: %d", total, 5)
	}
}

func TestDivisaoPor2(t *testing.T){
	total := DivisaoPor2(10)
	
	if total != 5 {
		t.Errorf("Resultado da Divisao é inválido: Resultado %d. Esperado: %d", total, 5)
	}
}