package main

import "fmt"

func main() {
	var matriz [3][3]int

	// Preenchendo a matriz
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matriz[i][j] = i + j
		}
	}

	// Imprimindo a matriz
	for i := 0; i < 3; i++ {
		fmt.Println(matriz[i])
	}

	// Criando uma matriz dinâmica 3x3 com slices
	matriz_slice := make([][]int, 3)
	for i := range matriz_slice {
		matriz_slice[i] = make([]int, 3)
	}

	// Preenchendo a matriz
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matriz_slice[i][j] = i + j
		}
	}

	// Imprimindo a matriz
	for _, linha := range matriz_slice {
		fmt.Println(linha)
	}

}
