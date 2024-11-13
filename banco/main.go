package main

import "fmt"

type ContaCorrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

func (cC *ContaCorrente) Sacar(valor float64) string { // onde cC *ContaCorrente equivale ao this, do Java
	valido := valor > 0 && valor <= cC.saldo
	if valido {
		cC.saldo -= valor
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (cC *ContaCorrente) Depositar(valor float64) string {
	valido := valor > 0
	if valido {
		cC.saldo += valor
		return "Deposito realizado com sucesso"
	} else {
		return "Operação inválida"
	}
}

func main() {
	var conta1 ContaCorrente
	conta1.titular = "Fhilype"
	conta1.agencia = 589
	conta1.conta = 123456
	conta1.saldo = 10000000.00
	fmt.Println(conta1)

	fmt.Println(conta1.Sacar(500), conta1)
	fmt.Println(conta1.Sacar(10000000.00), conta1)

	fmt.Println(conta1.Depositar(500), conta1)
	fmt.Println(conta1.Depositar(-500), conta1)
	fmt.Println(conta1.Sacar(1000.00), conta1)

	/*
		var conta1 *ContaCorrente
		conta1 = new(ContaCorrente)
		conta1.titular = "Fhilype"
		conta1.agencia = 589
		conta1.conta = 123456
		conta1.saldo = 10000000.00
		fmt.Println(conta1)
		fmt.Println(&conta1) // & mostra o endereço na memória

		var conta11 *ContaCorrente
		conta11 = new(ContaCorrente)
		conta11.titular = "Fhilype"
		conta11.agencia = 589
		conta11.conta = 123456
		conta11.saldo = 10000000.00
		fmt.Println(conta11)

		fmt.Println(conta1 == conta11) // false pois estão em endereços de memória diferentes
		fmt.Println(&conta11)

		fmt.Println(*conta1 == *conta11) // true pois seus valores são os mesmos

		conta2 := ContaCorrente{titular: "Marcia", agencia: 985, conta: 654321, saldo: 10000000.00}
		fmt.Println(conta2)

		fmt.Println(conta1 == conta2)
	*/
}
