package main

import "fmt"

// Função genérica simples
func identidade[T any /*pode ser comparable*/](valor T) T {
	return valor
}

// Restrição personalizada
type Numerico interface {
	int | float64
}

// Função genérica com restrição
func somar[T Numerico](a, b T) T {
	return a + b
}

func main() {
	fmt.Println("Identidade:")
	fmt.Println(identidade("Go"))
	fmt.Println(identidade(42))

	fmt.Println("\nSoma:")
	fmt.Println(somar(10, 5))    // int
	fmt.Println(somar(3.2, 4.8)) // float64
}
