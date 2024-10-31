package main

import "fmt"

func main() {
	idade := 20

	if idade >= 18 {
		fmt.Println("Você é maior de idade.")
	} else {
		fmt.Println("Você é menor de idade.")
	}

	nota := 85

	if nota >= 90 {
		fmt.Println("Nota A")
	} else if nota >= 80 {
		fmt.Println("Nota B")
	} else if nota >= 70 {
		fmt.Println("Nota C")
	} else {
		fmt.Println("Nota F")
	}
}
