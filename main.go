package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Chaitanya-Shahare/orders-microservice/application"
)

func main() {
	app := application.New();

	// derive a new context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)

	if err != nil {
		fmt.Println("failed to start app:", err)
	}

}