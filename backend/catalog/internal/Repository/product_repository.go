package repository

import "database/sql"

type PostGressProductRepository struct {
	db *sql.DB
}

func NewProductsRepository(db *sql.DB) *PostGressProductRepository {
	return &PostGressProductRepository{db: db}
}
