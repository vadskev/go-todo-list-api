package main

import (
	"context"
	"github.com/vadskev/go-todo-list-api/internal/app"
	"log"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("Failed to create app: %s", err.Error())
	}

	err = a.RunServer(ctx)
	if err != nil {
		log.Fatalf("Failed to run app: %s", err.Error())
	}
}
