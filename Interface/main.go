package main

import "fmt"

// Interface
type Animal interface {
	Falar() string
}

// Structs que implementam a interface
type Cachorro struct {
	Nome string
}

func (c Cachorro) Falar() string {
	return "Au au!"
}

type Gato struct {
	Nome string
}

func (g Gato) Falar() string {
	return "Miau!"
}

// Função que usa interface como parâmetro
func fazerAnimalFalar(a Animal) {
	fmt.Println(a.Falar())
}

func main() {
	c := Cachorro{Nome: "Rex"}
	g := Gato{Nome: "Mimi"}

	fazerAnimalFalar(c)
	fazerAnimalFalar(g)

	// Slice de interface
	animais := []Animal{c, g}
	for _, animal := range animais {
		fmt.Println(animal.Falar())
	}
}
