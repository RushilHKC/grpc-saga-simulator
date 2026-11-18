package main

import (
	"log"
	"net/http"
	"time"

	"github.com/RushilHKC/grpcSimulator/backend/Order/internal/config"
	"github.com/RushilHKC/grpcSimulator/backend/Order/internal/env"
	shared "github.com/RushilHKC/grpcSimulator/backend/shared/shutdown"
	database "github.com/RushilHKC/grpcSimulator/shared/gen/Db"
	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env file: ", err)
	}

	con := config.LoadDbConfig()

	db, err := database.New_db_connection(con.Port, con.UserName, con.Password, con.DbName, con.DbTimeOut)

	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)

	if err != nil {
		log.Fatal("Error Connecting DB:", err)
	} else {
		log.Print("DB connected successfully")
	}

	serverConfig := config.LoadServerConfig()
	handler := Handler.NewHandler()
	mux := http.NewServeMux()

	server := shared.CreateServer(serverConfig.Addr, mux)

	shutDownTime := env.GetInt("ShutDownTIme")
	shutDownTimeOut := time.Duration(shutDownTime) * time.Second

	if err := shared.RunServer(server, shutDownTimeOut); err != nil {
		log.Fatalf("Error starting Server:  %v", err)
	}
}
