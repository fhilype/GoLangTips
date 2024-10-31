package main

import "fmt"

// Definindo uma struct
type Pessoa struct {
    Nome  string
    Idade int
    Cidade string
}

// Método para a struct Pessoa
func (p Pessoa) Saudacao() {
    fmt.Printf("Oi, meu nome é %s e eu tenho %d anos. Moro em %s.\n", p.Nome, p.Idade, p.Cidade)
}

func main() {
    pessoa := Pessoa{Nome: "Alice", Idade: 30, Cidade: "São Paulo"}
    pessoa.Saudacao()
}
