package catalog_db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/env"
	_ "github.com/lib/pq"
)

func New_db_connection(port string, user, password, db_name string) (*sql.DB, error) {
	connection_str := fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable", port, user, password, db_name)

	db, err := sql.Open("postgres", connection_str)

	if err != nil {
		log.Fatal("Error connecting DB: ", err)
		return nil, err
	}

	db_time_out := env.GetInt("database_time_out")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(db_time_out)*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal("Error db connect timeOut: ", err)
		return nil, err
	}

	return db, nil

}
