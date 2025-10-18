package main

import "fmt"

const prefixoOlaPortugues = "Olá "
const prefixoOlaEspanhol = "Hola "
const prefixoOlafrances = "Bonjour "

func Ola(name string, lang string) string {
	if name == "" {
		name = "mundo"
	}

	if lang == "espanhol" {
		return prefixoOlaEspanhol + name
	}

	if lang == "francês" {
		return prefixoOlafrances + name
	}
	return prefixoOlaPortugues + name
}

func main() {
	fmt.Println(Ola("", ""))
}
