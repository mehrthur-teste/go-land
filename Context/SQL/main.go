package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // ou outro driver
)

func main() {
	db, err := sql.Open("postgres", "postgres://usuario:senha@localhost:5432/meubanco?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	query := "SELECT pg_sleep(2);" // simula uma query lenta
	_, err = db.ExecContext(ctx, query)

	if err != nil {
		fmt.Println("Erro:", err) // provavelmente será timeout
	} else {
		fmt.Println("Query executada com sucesso")
	}
}

// O código acima é um exemplo de como usar contextos com SQL em Go. Ele cria um contexto com timeout de 1 segundo e executa uma query que simula uma operação lenta. Se a operação demorar mais do que o timeout, um erro será retornado.
// O uso de contextos é importante para evitar que operações lentas bloqueiem o fluxo do programa, especialmente em aplicações que fazem chamadas a bancos de dados ou serviços externos. O código também demonstra como usar o pacote `database/sql` para interagir com um banco de dados PostgreSQL, mas o conceito se aplica a outros bancos de dados e drivers também.
// O código acima é um exemplo de como usar contextos com SQL em Go. Ele cria um contexto com timeout de 1 segundo e executa uma query que simula uma operação lenta. Se a operação demorar mais do que o timeout, um erro será retornado.
// O uso de contextos é importante para evitar que operações lentas bloqueiem o fluxo do programa, especialmente em aplicações que fazem chamadas a bancos de dados ou serviços externos. O código também demonstra como usar o pacote `database/sql` para interagir com um banco de dados PostgreSQL, mas o conceito se aplica a outros bancos de dados e drivers também.
// O código acima é um exemplo de como usar contextos com SQL em Go. Ele cria um contexto com timeout de 1 segundo e executa uma query que simula uma operação lenta. Se a operação demorar mais do que o timeout, um erro será retornado.
