package main

import "fmt"

// Definindo uma struct
type Pessoa struct {
    Nome  string
    Idade int
    Cidade string
}

func main() {
    // Criando uma instância de Pessoa
    pessoa := Pessoa{Nome: "Alice", Idade: 30, Cidade: "São Paulo"}

    // Acessando campos da struct
    fmt.Println("Nome:", pessoa.Nome)
    fmt.Println("Idade:", pessoa.Idade)
    fmt.Println("Cidade:", pessoa.Cidade)
}
