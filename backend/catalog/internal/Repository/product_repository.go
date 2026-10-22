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
	GetOrderDetails(ctx context.Context, id int) (*model.OrderDetails, error)
}

var _ ProductRepository = (*postgresProductRepository)(nil)

type postgresProductRepository struct {
	db *sql.DB
}

func NewProductsRepository(db *sql.DB) *postgresProductRepository {
	return &postgresProductRepository{db: db}
}

const getProductQuery = "SELECT id,name,description,price,currency FROM products WHERE id=$1"

func (r *postgresProductRepository) GetProduct(ctx context.Context, id int) (*model.Product, error) {

	var product model.Product

	err := r.db.QueryRowContext(ctx, getProductQuery, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Currency)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Error product not found: %s", err)
			return nil, err
		}

		return nil, err

	}

	return &product, nil
}

const getProductsQuery = "SELECT id,name,description,price,currency FROM products"

func (r *postgresProductRepository) GetProducts(ctx context.Context) ([]model.Product, error) {
	products := []model.Product{}

	rows, err := r.db.QueryContext(ctx, getProductsQuery)
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

const orderDetailsQuery = "SELECT id,price,currency FROM products WHERE id=$1"

func (r *postgresProductRepository) GetOrderDetails(ctx context.Context, id int) (*model.OrderDetails, error) {

	var details model.OrderDetails

	err := r.db.QueryRowContext(ctx, orderDetailsQuery, id).Scan(&details.ID, &details.Price, &details.Currency)

	if err != nil {
		log.Printf("Error product not found: %s", err)
		return nil, err
	}

	return &details, nil
}
