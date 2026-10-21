package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/model"
)

type ProductRepository interface {
	GetProduct(ctx context.Context, id int) (*model.Product, error)
	GetProducts(ctx context.Context) ([]model.Product, error)
}

var _ ProductRepository = (*postgresProductRepository)(nil)

type postgresProductRepository struct {
	db *sql.DB
}

func NewProductsRepository(db *sql.DB) *postgresProductRepository {
	return &postgresProductRepository{db: db}
}

func (r *postgresProductRepository) GetProduct(ctx context.Context, id int) (*model.Product, error) {

	var product model.Product

	query := "SELECT id,name,description,price,currency FROM products where id=$1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Currency)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Error product not found: %s", err)
			return nil, err
		}

		return nil, err

	}

	return &product, nil
}

func (r *postgresProductRepository) GetProducts(ctx context.Context) ([]model.Product, error) {
	products := []model.Product{}

	query := "SELECT id,name,description,price,currency FROM products"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		product := model.Product{}
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Currency); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil

}
