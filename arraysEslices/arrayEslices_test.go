package arrayseslices

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {
	numeros := []int{1, 2, 3, 4, 5}

	t.Run("colecao de 5 numeros", func(t *testing.T) {
		resultado := Soma(numeros)
		esperado := 15

		if esperado != resultado {
			t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
		}

	})
	t.Run("colecao de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3}

		resultado := Soma(numeros)
		esperado := 6

		if esperado != resultado {
			t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
		}
	})

}

func TestSomaTudo(t *testing.T) {

	resultado := SomaTudo([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("resultado %v esperado %v", resultado, esperado)
	}
}

func TestSomaTodoOResto(t *testing.T) {
	resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
	esperado := []int{2, 9}

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("resultado %v, esperado %v", resultado, esperado)
	}
}
