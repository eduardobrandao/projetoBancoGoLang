package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {

	contaBeatriz := contas.ContaCorrente{Titular: clientes.Titular{
		Nome: "Beatriz",
		CPF: "123.111.123.11",
		Profissao: "Enfermeira"},
	NumeroAgencia: 123, NumeroConta: 123456, Saldo: 1200}

	fmt.Println(contaBeatriz)
}

//contaEduardo := contas.ContaCorrente{Titular: "Eduardo", Saldo: 500}
	//contaCarlos := contas.ContaCorrente{Titular: "Carlos", Saldo: 800}

	//status := contaCarlos.Transferir(500, &contaEduardo)

	//fmt.Println(status)
	//fmt.Println(contaCarlos)
	//fmt.Println(contaEduardo)

	//contaEduardo := ContaCorrente{}
	//contaEduardo.titular = "Eduardo"
	//contaEduardo.saldo = 500

	//fmt.Println(contaEduardo.Sacar(300))
	//contaEduardo.Depositar(400)
	//fmt.Println(contaEduardo.saldo)
	//fmt.Println(contaEduardo.Depositar(400))
	//status, valor := contaEduardo.Depositar(400)
	//fmt.Println(status, valor)