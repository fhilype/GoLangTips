package contas

import "banco/clientes"

/*
	Visibilidade
	Atributos com a primeira letra em maiúsculo, significa que o atributo é público, do contrário, privado
*/
type ContaCorrente struct {
	Titular        clientes.Titular
	Agencia, Conta int
	saldo          float64
}

func (cC *ContaCorrente) Sacar(valor float64) string { // onde cC *ContaCorrente equivale ao this, do Java
	valido := valor > 0 && valor <= cC.saldo
	if valido {
		cC.saldo -= valor
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (cC *ContaCorrente) Depositar(valor float64) (string, float64) {
	valido := valor > 0
	if valido {
		cC.saldo += valor
		return "Deposito realizado com sucesso", cC.saldo
	} else {
		return "Operação inválida", cC.saldo
	}
}

func (cC *ContaCorrente) Transferir(valor float64, cCDestino *ContaCorrente) bool {
	valido := valor > 0 && cC.saldo >= valor
	if valido {
		cC.Sacar(valor)
		cCDestino.Depositar(valor)
		return true
	} else {
		return false
	}
}

func (cC *ContaCorrente) ObterSaldo() float64 {
	return cC.saldo
}
