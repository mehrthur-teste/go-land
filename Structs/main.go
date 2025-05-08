package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) Saudacao() {
	fmt.Printf("Olá! Sou %s e tenho %d anos.\n", p.Nome, p.Idade)
}

func (p *Pessoa) Envelhecer() {
	p.Idade++
}

func main() {
	p := Pessoa{Nome: "Samir", Idade: 29}

	p.Saudacao() // Olá! Sou Samir e tenho 29 anos.
	p.Envelhecer()
	p.Saudacao() // Olá! Sou Samir e tenho 30 anos.
}
