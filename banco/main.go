package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func PagarBoleto(c Conta, valor float64) {
	c.Sacar(valor)
}

type Conta interface {
	Sacar(valor float64) string
}

func main() {
	cliente1 := clientes.Titular{Nome: "Fhilype", CPF: "999.999.999-99", Profissao: "Engenheiro"}
	conta1 := contas.ContaCorrente{Titular: cliente1, Agencia: 589, Conta: 123456}
	conta1.Depositar(10000.00)

	cliente2 := clientes.Titular{Nome: "Marcia", CPF: "888.888.888-88", Profissao: "Advogada"}
	conta2 := contas.ContaCorrente{Titular: cliente2, Agencia: 985, Conta: 654321}
	conta2.Depositar(10000.00)
	PagarBoleto(&conta2, 1000.00)

	conta3 := contas.ContaPoupanca{Titular: cliente1, Agencia: 589, Conta: 123456, Operacao: 1}
	conta3.Depositar(30000.00)
	PagarBoleto(&conta3, 700.00)

	conta1.Transferir(1000, &conta2)

	fmt.Println(conta1.ObterSaldo())
	fmt.Println(conta2.ObterSaldo())
	fmt.Println(conta3.ObterSaldo())

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
