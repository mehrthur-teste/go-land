package main

import "fmt"

func main() {

	// Mapa aninhado
	users := map[string]map[string]string{
		"user1": {
			"name":  "Alice",
			"email": "alice@example.com",
		},
		"user2": {
			"name":  "Bob",
			"email": "bob@example.com",
		},
	}

	// Acessando um valor específico
	fmt.Println("Email do user1:", users["user1"]["email"])

	// Iterando sobre o mapa
	for id, info := range users {
		fmt.Printf("ID: %s\n", id)
		for key, value := range info {
			fmt.Printf("  %s: %s\n", key, value)
		}
	}

	// Adicionando um novo user
	users["user3"] = map[string]string{
		"name":  "Carol",
		"email": "carol@example.com",
	}

	users_other_example := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}

	users_other_example["Charlie"] = 40

	for name, age := range users_other_example {
		fmt.Printf("%s tem %d anos\n", name, age)
	}

}
