package main

import "testing"

func TestOla(t *testing.T) {
	verificarMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper() // diz para o conjunto de testes que este e um metodo auxiliar
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}
	//t.Run sao subtestes para testar cenarios diferentes em uma funcao
	t.Run("diz ola para as pessoas", func(t *testing.T) {
		resultado := Ola("Carlos", "")
		esperado := "Olá Carlos"
		verificarMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'Olá mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá mundo"
		verificarMensagemCorreta(t, resultado, esperado)

	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Eloide", "espanhol")
		esperado := "Hola Eloide"
		verificarMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em frances", func(t *testing.T) {
		resultado := Ola("Eloide", "francês")
		esperado := "Bonjour Eloide"
		verificarMensagemCorreta(t, resultado, esperado)
	})

}
