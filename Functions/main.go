package main

import "fmt"

// Função simples sem parâmetros
func saudacao() {
	fmt.Println("Olá, mundo!")
}

// Função com parâmetros e retorno
func somar(a, b int) int {
	return a + b
}

// Função com múltiplos retornos
func dividir(a, b int) (int, int) {
	quociente := a / b
	resto := a % b
	return quociente, resto
}

// Função que recebe um slice
func imprimirValores(valores []int) {
	for _, v := range valores {
		fmt.Println("Valor:", v)
	}
}

// Função que cria e preenche uma matriz 3x3
func criarMatriz3x3() [][]int {
	matriz := make([][]int, 3)
	for i := range matriz {
		matriz[i] = make([]int, 3)
		for j := range matriz[i] {
			matriz[i][j] = i + j
		}
	}
	return matriz
}

func main() {
	saudacao()

	resultado := somar(5, 7)
	fmt.Println("Soma:", resultado)

	q, r := dividir(10, 3)
	fmt.Printf("Divisão: quociente = %d, resto = %d\n", q, r)

	numeros := []int{10, 20, 30}
	fmt.Println("Imprimindo slice:")
	imprimirValores(numeros)

	fmt.Println("Matriz 3x3:")
	matriz := criarMatriz3x3()
	for _, linha := range matriz {
		fmt.Println(linha)
	}

	// Função anônima
	f := func(nome string) {
		fmt.Println("Olá,", nome)
	}
	f("Samir")
}
