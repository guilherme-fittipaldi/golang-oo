package main

import( 
	"fmt"
	c "github.com/guilherme-fittipaldi/golang-oo/contas"
)

func main() {

	contaDaSilvia := c.ContaCorrente{}
	contaDaSilvia.Titular = "Silvia"
	contaDaSilvia.Saldo = 300
	contaDoGustavo := c.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	status := contaDoGustavo.Transferir(-200, &contaDaSilvia)
	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

	// contaDoGuilherme := ContaCorrente{titular: "Guilherme",
	// 	numeroAgencia: 589, numeroConta: 123456, Saldo: 125.5}

	// contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}
	// fmt.Println(contaDoGuilherme)
	// fmt.Println(contaDaBruna)

	// var contaDaCris *ContaCorrente
	// contaDaCris = new(ContaCorrente)
	// contaDaCris.titular = "Cris"
	// contaDaCris.saldo = 500

	// fmt.Println(*contaDaCris)
}
