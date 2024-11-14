package contas

/*
	Visibilidade
	Atributos com a primeira letra em maiúsculo, significa que o atributo é público, do contrário, privado
*/
type ContaCorrente struct {
	Titular string
	Agencia int
	Conta   int
	Saldo   float64
}

func (cC *ContaCorrente) Sacar(valor float64) string { // onde cC *ContaCorrente equivale ao this, do Java
	valido := valor > 0 && valor <= cC.Saldo
	if valido {
		cC.Saldo -= valor
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (cC *ContaCorrente) Depositar(valor float64) (string, float64) {
	valido := valor > 0
	if valido {
		cC.Saldo += valor
		return "Deposito realizado com sucesso", cC.Saldo
	} else {
		return "Operação inválida", cC.Saldo
	}
}

func (cC *ContaCorrente) Transferir(valor float64, cCDestino *ContaCorrente) bool {
	valido := valor > 0 && cC.Saldo >= valor
	if valido {
		cC.Sacar(valor)
		cCDestino.Depositar(valor)
		return true
	} else {
		return false
	}
}
