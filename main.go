package main

import "fmt"

type ContaCorrente struct {
	titular string
	numeroAgencia int
	numeroConta int
	saldo float64
} 

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {

	contaEduardo := ContaCorrente{}
	contaEduardo.titular = "Eduardo"
	contaEduardo.saldo = 500

	fmt.Println(contaEduardo.saldo)
	 
	fmt.Println(contaEduardo.Sacar(300))

	fmt.Println(contaEduardo.saldo)
}