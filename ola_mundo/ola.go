package main

import "fmt"

const prefixoOlaPortugues = "Olá "
const prefixoOlaEspanhol = "Hola "
const prefixoOlafrances = "Bonjour "

func Ola(name string, lang string) string {
	if name == "" {
		name = "mundo"
	}

	return prefixodeSaudacao(lang) + name

}

func prefixodeSaudacao(lang string) (prefixo string) {

	switch lang {

	case "frances":
		prefixo = prefixoOlafrances
	case "espanhol":
		prefixo = prefixoOlaEspanhol
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}
func main() {
	fmt.Println(Ola("", ""))
}
