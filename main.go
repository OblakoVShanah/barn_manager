package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/OblakoVShanah/barn_manager/internal/app"
)

// main is the entry point of the application.
// It initializes the application configuration, sets up the database connection,
// and starts the HTTP server with graceful shutdown handling.
func main() {
	ctx := context.Background()

	config, err := app.NewConfig("configs/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	app, err := app.New(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	if err = app.Setup(ctx, config.DB.DSN); err != nil {
		log.Fatal(err)
	}

	// Set up graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start the application in a goroutine
	go func() {
		if err = app.Start(); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-sigChan
	log.Println("Shutdown signal received, exiting...")
}
