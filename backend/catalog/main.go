package main

import (
	"context"
	"log"
	"net/http"
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

func runServer(server *http.Server) error {

	serverErr := make(chan error, 1)

	// go func() {
	if err := server.ListenAndServe(); err != nil {
		serverErr <- err
	}
	// }()

	select {
	case err := <-serverErr:
		return err
	default:
		return nil
	}

}

func main() {

	server := createServer()
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	if err := runServer(server); err != nil {
		log.Fatalf("Error starting Server %v", err)
	}

}
