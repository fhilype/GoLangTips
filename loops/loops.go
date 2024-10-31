package main

import "fmt"

func main() {
	// Loop simples
	for i := 0; i < 5; i++ {
		fmt.Println("Contagem: ", i)
	}

	// Loop tipo while
	j := 0
	for j < 5 {
		fmt.Println("Enquanto: ", j)
		j++
	}

	// Loop infinito
	k := 0
	for {
		fmt.Println("Infinito: ", k)
		k++
		if k == 5 {
			break
		}
	}
}