package main

import (
	"log"
	"net/http"
	"time"

	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/catalog_db"
	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/config"
	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/env"
	handlerSVC "github.com/RushilHKC/grpcSimulator/backend/catalog/internal/handler"
	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/repository"
	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/routes"
	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/service"
	"github.com/RushilHKC/grpcSimulator/backend/shared/shutdown/shared"
	"github.com/joho/godotenv"
)

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
	mux := http.NewServeMux()

	routes.Register(mux, handler)
	server := shared.CreateServer(serverConfig.Addr, mux)

	shut_down_time := env.GetInt("shut_down_time_out")
	shutDownTimeOut := time.Duration(shut_down_time) * time.Second

	if err := shared.RunServer(server, shutDownTimeOut); err != nil {
		log.Fatalf("Error starting Server:  %v", err)
	}

}
