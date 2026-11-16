package main

import (
	"log"
	"net/http"
	"time"

	"github.com/RushilHKC/grpcSimulator/backend/Order/internal/config"
	shared "github.com/RushilHKC/grpcSimulator/backend/shared/shutdown"
	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env file: ", err)
	}

	con := config.loadDbConfig()

	db, err := Db.NewDbConnection(con.Port, con.user, con.Password, con.DbName)

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
