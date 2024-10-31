package main

import "fmt"

func main() {
    // Criando um map
    idade := map[string]int{
        "Alice": 30,
        "Bob": 25,
    }
    
    // Acessando valores
    fmt.Println("Idade da Alice:", idade["Alice"])
    
    // Adicionando um novo par chave-valor
    idade["Charlie"] = 20
    fmt.Println("Idade atualizada:", idade)
    
    // Removendo um par chave-valor
    delete(idade, "Bob")
    fmt.Println("Idade após remoção:", idade)
}
