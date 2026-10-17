package catalog_db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func New_db_connection(port string, user, password, db_name string) (*sql.DB, error) {
	connection_str := fmt.Sprintf("port=%s user=%s password=%s dbname=%s sslmode=disable", port, user, password, db_name)

	db, err := sql.Open("postgres", connection_str)

	if err != nil {
		log.Fatal("Error connecting DB:", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5000)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal("Error db connect timeOut", err)
		return nil, err
	}

	return db, nil

}
