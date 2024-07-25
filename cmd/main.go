package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/vadskev/go_final_project/internal/app"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	a, err := app.NewApp()
	if err != nil {
		log.Fatalf("Failed to create app: %s", err.Error())
	}

	err = a.RunServer(ctx)
	if err != nil {
		log.Fatalf("Failed to run app: %s", err.Error())
	}
}
