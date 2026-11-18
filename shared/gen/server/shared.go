package shared

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func createServer(addr string, mux http.Handler) *http.Server {

	// mux := http.NewServeMux()

	// mux.HandleFunc("GET /products", handler.GetProducts)
	// mux.HandleFunc("GET /product/{id}", handler.GetProduct)

	return &http.Server{
		Addr:    addr,
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
		fmt.Println("Error starting server: ", err)
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
