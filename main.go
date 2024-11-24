package main

import (
	"chirpify/config"
	"chirpify/helper"
	"chirpify/router"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Base context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Load configuration
	config, err := config.LoadConfig("config/app.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Get port from configuration
	port := config.Server.Port
	if port == 0 {
		log.Fatal("Port not specified in configuration")
	}

	// Initialize dependencies
	deps := router.InitializeDependencies()

	//router
	routes := router.NewRouter(deps)

	server := &http.Server{
		Addr:    ":" + helper.IntToString(port),
		Handler: routes,
	}

	done := make(chan bool)

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("HTTP server ListenAndServe: %v", err)
			helper.ErrorPanic(err)
		}
		close(done)
	}()

	// Listening for shutdown signal in a goroutine
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		// Shutdown signal received, cancel the context
		log.Println("Signal received, initiating graceful shutdown")
		cancel()

		// Graceful shutdown
		shutdownCtx, cancelShutdown := context.WithTimeout(ctx, 30*time.Second)
		defer cancelShutdown()
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
			helper.ErrorPanic(err)
		}

		close(done)
	}()

	// Wait here until the done channel is closed
	<-done
	log.Println("Server stopped")
}
