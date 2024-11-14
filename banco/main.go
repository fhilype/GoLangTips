package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	conta1 := contas.ContaCorrente{Titular: "Fhilype", Agencia: 589, Conta: 123456, Saldo: 10000.00}

	conta2 := contas.ContaCorrente{Titular: "Marcia", Agencia: 985, Conta: 654321, Saldo: 10000.00}

	conta1.Transferir(1000, &conta2)

	fmt.Println(conta1)
	fmt.Println(conta2)

	/*
		Para obter os demais retornos de uma função com múltiplos retornos
		mensagem, saldo := conta1.Depositar(500)
		fmt.Println(mensagem, saldo)
		mensagem, saldo = conta1.Depositar(-500)
		fmt.Println(mensagem, saldo)
	*/

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
