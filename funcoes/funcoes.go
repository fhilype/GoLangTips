package main

import "fmt"

// Função simples que não recebe parâmetros e não retorna valores
func saudacao() {
    fmt.Println("Olá, bem-vindo às funções em Go!")
}

// Função que recebe parâmetros
func soma(a int, b int) int {
    return a + b
}

// Função com múltiplos retornos
func divide(a int, b int) (int, int) {
    quociente := a / b
    resto := a % b
    return quociente, resto
}

func main() {
    saudacao()

    resultado := soma(3, 4)
    fmt.Println("Resultado da soma:", resultado)

    quociente, resto := divide(10, 3)
    fmt.Println("Quociente:", quociente, "Resto:", resto)
}
