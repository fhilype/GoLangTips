package main

import (
	"fmt"
	"meu_projeto/pessoa"
)

func main() {
	alice := pessoa.Pessoa{Nome: "Alice", Idade: 30, Cidade: "São Paulo"}
	fmt.Println("Idade antes do aniversário:", alice.Idade)

	alice.FazerAniversario()

	fmt.Println("Idade depois do aniversário:", alice.Idade)
}
