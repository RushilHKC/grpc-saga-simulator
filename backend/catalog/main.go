package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	handler "github.com/RushilHKC/grpcSimulator/backend/catalog/internal/handler"
	"github.com/joho/godotenv"
)

// Endpoint	Purpose
// GET /products	List products
// GET /products/{id}	Fetch one product
// GET /products?category=electronics	Filter by category
// GET /products?limit=20&offset=0	Paginate results
// GET /products?sort=price_asc	Sort by price

type config struct {
	addr string
}

func createServer(cfg config) *http.Server {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handler.GetProducts)
	mux.HandleFunc("GET /{id}", handler.GetProduct)
	// mux.HandleFunc("GET /{category}", handler.FilterByCategory)

	return &http.Server{
		Addr:    cfg.addr,
		Handler: mux,
	}
}

func runServer(server *http.Server, shutdownTimeOut time.Duration) error {

	serverErr := make(chan error, 1)

	go func() {
		fmt.Println("starting server")
		if err := server.ListenAndServe(); err != nil {
			serverErr <- err
		}
		close(serverErr)
	}()

	stop := make(chan os.Signal, 1)
	defer close(stop)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		fmt.Println("Error starting server")
		return err
	// case <-ctx.Done():
	// 	fmt.Println("Context cancelled")
	case <-stop:
		fmt.Println("Shutdown signal received")
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

	addr, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("Missing port in env")
	}

	cfg := config{
		addr: addr,
	}

	server := createServer(cfg)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	// ctx := context.Background()

	// ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	// defer cancel()

	shutDownTimeOut := time.Second * 10

	if err := runServer(server, shutDownTimeOut); err != nil {
		log.Fatalf("Error starting Server %v", err)
	}

}
