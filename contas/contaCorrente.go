package contas

import "github.com/guilherme-fittipaldi/golang-oo/clientes"

type ContaCorrente struct {
	Titular       string
	numeroAgencia int
	numeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.Saldo
	} else {
		return "Valor do depósito menor que zero", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia > 0 {
		if c.Saldo > valorDaTransferencia {
			c.Saldo -= valorDaTransferencia

			contaDestino.Depositar(valorDaTransferencia)
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
