package main

import "fmt"

func main() {
	// Criando uma slice de inteiros
	numeros := []int{1, 2, 3, 4, 5}

	// Acessando elementos
	fmt.Println("Primeiro número:", numeros[0])

	// Adicionando elementos
	numeros = append(numeros, 6)
	fmt.Println("Slice com novo número:", numeros)

	// Subslice
	parte := numeros[1:4]
	fmt.Println("Subslice:", parte)
}
