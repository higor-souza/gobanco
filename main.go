package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	contaDoDenis.Sacar(50000)

	fmt.Println(contaDoDenis.ObterSaldo())
	PagarBoleto(&contaDoDenis, 60)
	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaPati := contas.ContaCorrente{}
	fmt.Println(contaDaPati)
	contaDaPati.Depositar(100)
	PagarBoleto(&contaDaPati, 10)
	fmt.Println(contaDaPati)

	/* contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)

	fmt.Println(contaExemplo.ObterSaldo()) */

	// contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	// contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	// fmt.Println(contaDaSilvia, contaDoGustavo)

	// status := contaDaSilvia.Transferir(200, &contaDoGustavo)
	// fmt.Println(status)

	// fmt.Println(contaDaSilvia, contaDoGustavo)

	// contaDaSilvia := ContaCorrente{}
	// contaDaSilvia.titular = "Silvia"
	// contaDaSilvia.saldo = 500

	//fmt.Println(contaDaSilvia.saldo)

	//fmt.Println(contaDaSilvia.Sacar(400))
	//fmt.Println(contaDaSilvia.saldo)

	//fmt.Println(contaDaSilvia.saldo)
	//contaDaSilvia.Depositar(500)
	//status, valor := contaDaSilvia.Depositar(2000)
	//fmt.Println(status, valor)

}
