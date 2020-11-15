package main

import "testing"

func TestGreating(t *testing.T) {
	greating := greating("Code.education Rocks!")
	if greating != "<b>Code.education Rocks!</b>" {
		t.Errorf("Resposta errada. O Retorno foi: %s e deveria retornar: %s.", greating, "<b>Code.education Rocks!</b>");
	}
}