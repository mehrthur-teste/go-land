package main

import "fmt"

func main() {
	// Criando um slice
	numeros := []int{5, 10, 15}
	fmt.Println("Inicial:", numeros)

	// Adicionando elementos
	numeros = append(numeros, 20, 25)
	fmt.Println("Após append:", numeros)

	// Fatiando
	parte := numeros[1:4]
	fmt.Println("Slice parcial:", parte)

	// Iterando
	for i, v := range numeros {
		fmt.Printf("Posição %d: %d\n", i, v)
	}

	nums := []int{10, 20, 30, 40, 50}
	fmt.Println(nums[1:4]) // [20 30 40]
	fmt.Println(nums[:3])  // [10 20 30]
	fmt.Println(nums[2:])  // [30 40 50]
}
