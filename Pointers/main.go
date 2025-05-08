package main

import "fmt"

type Conta struct {
	Saldo float64
}

func (c *Conta) Depositar(valor float64) {
	c.Saldo += valor
}

func main() {
	conta := Conta{Saldo: 100}
	conta.Depositar(50)
	fmt.Println("Saldo:", conta.Saldo) // 150

	x := 10
	y := &x
	fmt.Println("Antes:", x)
	*y = 99
	fmt.Println("Depois:", x) // 99
}
