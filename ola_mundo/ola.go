package main

import "fmt"

const prefixoOlaPortugues = "Olá "
const prefixoOlaEspanhol = "Hola "
const prefixoOlafrances = "Bonjour "

func Ola(name string, lang string) string {
	if name == "" {
		name = "mundo"
	}

	prefixo := prefixoOlaPortugues

	switch lang {
	case "frances":
		prefixo = prefixoOlafrances
	case "espanhol":
		prefixo = prefixoOlaEspanhol
	}

	return prefixo + name
}

func main() {
	fmt.Println(Ola("", ""))
}
