package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

// Método com receiver por valor (não altera o original)
func (p Pessoa) Saudacao() {
	fmt.Printf("Oi, eu sou %s e tenho %d anos.\n", p.Nome, p.Idade)
}

// Método com receiver por ponteiro (altera o original)
func (p *Pessoa) Fazer() {
	p.Idade++
}

func main() {
	p := Pessoa{Nome: "João", Idade: 30}

	p.Saudacao() // Oi, eu sou João e tenho 30 anos.
	p.Fazer()    // incrementa idade
	p.Saudacao() // Oi, eu sou João e tenho 31 anos.
}
