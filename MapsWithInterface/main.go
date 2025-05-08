package main

import "fmt"

func main() {
	users := map[string]map[string]interface{}{
		"user1": {
			"name":    "Alice",
			"email":   "alice@example.com",
			"premium": true,
			"age":     30,
		},
		"user2": {
			"name":    "Bob",
			"email":   "bob@example.com",
			"premium": false,
			"age":     25,
		},
	}

	fmt.Println("Email do user1:", users["user1"]["email"])

	for id, info := range users {
		fmt.Printf("ID: %s\n", id)
		for key, value := range info {
			fmt.Printf("  %s: %v\n", key, value)
		}
	}
}
