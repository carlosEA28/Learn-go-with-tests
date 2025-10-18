package iteracao

import "fmt"

func ExampleRepetir() {
	letras := Repetir(10, "a")
	fmt.Println(letras)
	// Output: aaaaaaaaaa
}

func Repetir(quantidadeRepeticoes int, caracter string) string {
	var qtd = quantidadeRepeticoes
	var repeticoes string

	for i := 0; i < qtd; i++ {
		repeticoes += caracter
	}

	return repeticoes
}
