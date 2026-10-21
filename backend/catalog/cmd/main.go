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

	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/catalog_db"
	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/config"
	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/env"
	handlerSVC "github.com/RushilHKC/grpcSimulator/backend/catalog/internal/handler"
	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/repository"
	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/service"
	"github.com/joho/godotenv"
)

func createServer(cfg *config.Server_config, handler handlerSVC.Handler) *http.Server {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /products", handler.GetProducts)
	mux.HandleFunc("GET /product/{id}", handler.GetProduct)

	return &http.Server{
		Addr:    cfg.Addr,
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

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	//Load Environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file", err)
	}

	//load Database configurations
	con := config.Load_Db_Config()
	db, err := catalog_db.New_db_connection(con.Port, con.User, con.Password, con.Db_name)
	if err != nil {
		log.Fatal("Error connecting DB")

	} else {
		log.Print("DB connected Successfully")
	}

	productsRepo := repository.NewProductsRepository(db)
	productService := service.NewProductService(productsRepo)

	serverConfig := config.LoadServerConfig()
	handler := handlerSVC.NewHandler(productService)

	server := createServer(serverConfig, handler)

	shut_down_time := env.GetInt("shut_down_time_out")
	shutDownTimeOut := time.Duration(shut_down_time) * time.Second

	if err := runServer(server, shutDownTimeOut); err != nil {
		log.Fatalf("Error starting Server %v", err)
	}

}
