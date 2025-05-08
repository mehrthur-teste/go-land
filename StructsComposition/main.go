package main

import "fmt"

// Struct composta nomeada
type Endereco struct {
	Rua    string
	Cidade string
}

// Struct com composição nomeada
type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco // acesso via p.Endereco.Cidade
}

// Struct embutida (composição implícita)
type Animal struct{}

func (a Animal) Falar() {
	fmt.Println("Animal faz barulho")
}

type Cachorro struct {
	Animal // embedded: herda Falar()
	Nome   string
}

func main() {
	// Exemplo com composição nomeada
	p := Pessoa{
		Nome:  "Maria",
		Idade: 28,
		Endereco: Endereco{
			Rua:    "Rua das Palmeiras",
			Cidade: "Belo Horizonte",
		},
	}

	fmt.Println("Pessoa:")
	fmt.Println("Nome:", p.Nome)
	fmt.Println("Cidade:", p.Endereco.Cidade)

	// Exemplo com composição embutida
	c := Cachorro{Nome: "Rex"}

	fmt.Println("\nCachorro:")
	fmt.Println("Nome:", c.Nome)
	c.Falar() // Método herdado via embedding
}
