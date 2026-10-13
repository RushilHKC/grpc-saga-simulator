package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	handler "github.com/RushilHKC/grpcSimulator/backend/catalog/internal"
)

// Endpoint	Purpose
// GET /products	List products
// GET /products/{id}	Fetch one product
// GET /products?category=electronics	Filter by category
// GET /products?limit=20&offset=0	Paginate results
// GET /products?sort=price_asc	Sort by price

func createServer() *http.Server {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handler.GetProducts)
	mux.HandleFunc("GET /{id}", handler.GetProduct)
	// mux.HandleFunc("GET /{category}", handler.FilterByCategory)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}

func runServer(ctx context.Context, server *http.Server, shutdownTimeOut time.Duration) error {

	serverErr := make(chan error, 1)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			serverErr <- err
		}
	}()

	stop := make(chan os.Signal)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		return err
	case <-ctx.Done():
		log.Println("Context cancelled")
	case <-stop:
		log.Println("Shutdown signal received")
	}

	shutDownContext, cancel := context.WithTimeout(
		context.Background(),
		shutdownTimeOut,
	)

	defer cancel()

	if shutdownErr := server.Shutdown(shutDownContext); shutdownErr != nil {
		if CloseErr := server.Close(); CloseErr != nil {
			return errors.Join(CloseErr, shutdownErr)
		}

		return shutdownErr
	}

	return nil

}

func main() {

	server := createServer()
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	shutDownTimeOut := time.Second * 10

	if err := runServer(ctx, server, shutDownTimeOut); err != nil {
		log.Fatalf("Error starting Server %v", err)
	}

}
