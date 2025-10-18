package iteracao

import "testing"

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir(10, "a")
	}
}

func TestIteracao(t *testing.T) {
	repeticoes := Repetir(10, "a")
	esperado := "aaaaaaaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticoes)
	}
}
