package main

import (
	"fmt"
)

func main() {

	var numbers [5]int

	// Initialize the array with values
	for i := 0; i <= 4; i++ {
		numbers[i] = i * 2
		fmt.Printf("Index %d: %d\n", i, numbers[i])
	}

	//print the array one way
	fmt.Println("Array values:")
	fmt.Println(numbers)

	//print the array another way range numbers
	for i, value := range numbers {
		fmt.Printf("Index %d: %d\n", i, value)
	}

	// Criando um array com inferência de tamanho
	pares := [...]int{2, 4, 6, 8, 10, 11}
	fmt.Println("\nArray de pares:")
	for _, v := range pares {
		fmt.Println(v)
	}
	// Criando um slice a partir do array
	index := pares[:]
	fmt.Println("\nSlice de pares:")
	for _, v := range index {
		fmt.Println(v)
	}

	// Comparando arrays
	a := [3]string{"Go", "é", "legal"}
	b := [3]string{"Goes", "é", "legal"}

	fmt.Println("\nArrays iguais?", a == b)

}
