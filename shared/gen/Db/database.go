package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func New_db_connection(port, user, password, db_name string, db_time_out int) (*sql.DB, error) {
	connection_str := fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable", port, user, password, db_name)

	db, err := sql.Open("postgres", connection_str)

	if err != nil {
		log.Println("Error connecting DB: ", err)
		return nil, err
	}

	// db_time_out := env.GetInt("database_time_out")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(db_time_out)*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Println("Error db connect timeOut: ", err)
		return nil, err
	}

	return db, nil

}
