package main

import (
	"context"
	"fmt"
	"time"
)

func operacaoDemorada(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second):
		return fmt.Errorf("terminou depois de 3s")
	case <-ctx.Done():
		return ctx.Err() // retorna context deadline exceeded
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := operacaoDemorada(ctx)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Operação concluída com sucesso!")
	}
}

// O código acima cria um contexto com um timeout de 2 segundos e executa uma operação demorada que leva 3 segundos para ser concluída. Como o timeout é menor que o tempo de execução da operação, o erro "context deadline exceeded" será retornado.
// O uso de contextos é uma prática comum em Go para gerenciar cancelamentos e deadlines de operações assíncronas, especialmente em aplicações que fazem chamadas de rede ou operações que podem demorar muito tempo.
// O código acima demonstra o uso de contextos em Go, especificamente o uso de `context.WithTimeout` para definir um limite de tempo para uma operação. A função `operacaoDemorada` simula uma operação que leva 3 segundos para ser concluída, enquanto o contexto tem um timeout de 2 segundos. Quando a operação demora mais do que o tempo definido no contexto, um erro é retornado indicando que o prazo foi excedido.
