package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Cria um contexto com timeout de 2 segundos
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	done := make(chan string)

	go func() {
		time.Sleep(3 * time.Second) // simula operação lenta
		done <- "Resposta pronta!"
	}()

	select {
	case <-ctx.Done():
		http.Error(w, "Timeout! A operação demorou demais.", http.StatusRequestTimeout)
	case result := <-done:
		fmt.Fprintln(w, result)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
