package main

import (
	"context"
	"fmt"
)

func logRequest(ctx context.Context) {
	requestID := ctx.Value("request_id")
	fmt.Println("ID da requisição:", requestID)
}

func main() {
	ctx := context.WithValue(context.Background(), "request_id", "abc123")
	logRequest(ctx)
}
