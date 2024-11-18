package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular                  clientes.Titular
	Agencia, Conta, Operacao int
	saldo                    float64
}

func (cC *ContaPoupanca) Sacar(valor float64) string { // onde cC *ContaPoupanca equivale ao this, do Java
	valido := valor > 0 && valor <= cC.saldo
	if valido {
		cC.saldo -= valor
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (cC *ContaPoupanca) Depositar(valor float64) (string, float64) {
	valido := valor > 0
	if valido {
		cC.saldo += valor
		return "Deposito realizado com sucesso", cC.saldo
	} else {
		return "Operação inválida", cC.saldo
	}
}

func (cC *ContaPoupanca) ObterSaldo() float64 {
	return cC.saldo
}
