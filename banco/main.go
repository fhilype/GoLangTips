package main

import "fmt"

type ContaCorrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

func main() {
	var conta1 ContaCorrente
	conta1.titular = "Fhilype"
	conta1.agencia = 589
	conta1.conta = 123456
	conta1.saldo = 10000000.00

	fmt.Println(conta1)

	conta2 := ContaCorrente{titular: "Marcia", agencia: 985, conta: 654321, saldo: 10000000.00}

	fmt.Println(conta2)
}
