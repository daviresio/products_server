package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"products_server/config"
	"products_server/server"
	"syscall"
	"time"
)

func main() {
	cfg, _ := config.LoadConfig()

	dbPool, err := config.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v\n", err)
	}

	defer dbPool.Close()

	srv := server.NewServer(dbPool)
	go func() {
		log.Println("Starting server on :8080")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Failed to start server: %v\n", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown server: %v\n", err)
	}

	log.Println("Server stopped")
}
