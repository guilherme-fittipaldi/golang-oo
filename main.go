package main

import( 
	"fmt"
	"github.com/guilherme-fittipaldi/golang-oo/contas"
)


func main() {

	contaDaSilvia := contas.ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 300
	contaDoGustavo := contas.ContaCorrente{titular: "Gustavo", saldo: 100}

	status := contaDoGustavo.Transferir(-200, &contaDaSilvia)
	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

	// contaDoGuilherme := ContaCorrente{titular: "Guilherme",
	// 	numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	// contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}
	// fmt.Println(contaDoGuilherme)
	// fmt.Println(contaDaBruna)

	// var contaDaCris *ContaCorrente
	// contaDaCris = new(ContaCorrente)
	// contaDaCris.titular = "Cris"
	// contaDaCris.saldo = 500

	// fmt.Println(*contaDaCris)
}
