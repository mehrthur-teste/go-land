package main

import "fmt"

// Closure que mantém estado interno
func contador() func() int {
	contadorInterno := 0
	return func() int {
		contadorInterno++
		return contadorInterno
	}
}

// Closure com parâmetro externo capturado
func multiplicador(fator int) func(int) int {
	return func(x int) int {
		return x * fator
	}
}

func main() {
	fmt.Println("=== Closure contador ===")
	conta := contador()

	fmt.Println(conta()) // 1
	fmt.Println(conta()) // 2
	fmt.Println(conta()) // 3

	fmt.Println("\nOutro contador independente:")
	novoConta := contador()
	fmt.Println(novoConta()) // 1

	fmt.Println("\n=== Closure multiplicador ===")
	multiplicaPor2 := multiplicador(2)
	multiplicaPor3 := multiplicador(3)

	fmt.Println("2 x 5 =", multiplicaPor2(5)) // 10
	fmt.Println("3 x 5 =", multiplicaPor3(5)) // 15
}
