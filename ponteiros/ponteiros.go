package main

import "fmt"

type Pessoa struct {
    Nome  string
    Idade int
    Cidade string
}

// Método que modifica a struct usando um ponteiro
func (p *Pessoa) FazerAniversario() {
    p.Idade++
}

func main() {
    pessoa := Pessoa{Nome: "Alice", Idade: 30, Cidade: "São Paulo"}
    fmt.Println("Idade antes do aniversário:", pessoa.Idade)
    
    pessoa.FazerAniversario()
    
    fmt.Println("Idade depois do aniversário:", pessoa.Idade)
}
