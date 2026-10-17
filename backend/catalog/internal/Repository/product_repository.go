package repository

import (
	"database/sql"
	"log"

	p "github.com/RushilHKC/grpcSimulator/backend/catalog/internal/model"
)

type ProductRepository interface {
	GetProduct(id int) (*p.Product, error)
	GetProducts() ([]p.Product, error)
}

// var _ ProductRepository = (*postgresProductRepository)(nil)


type postgresProductRepository struct {
	db *sql.DB
}

func NewProductsRepository(db *sql.DB) *postgresProductRepository {
	return &postgresProductRepository{db: db}
}

func (r *postgresProductRepository) GetProduct(id int) (*p.Product, error) {

	var product p.Product

	query := "SELECT * FROM products where id = $1"
	err := r.db.QueryRow(query, id).Scan(product.ID, product.Currency, product.Description, product.Name, product.Price)

	if err != nil {

		if err == sql.ErrNoRows {
			log.Printf("Error product not found")
			return nil, err
		}

		return nil, err

	}

	return &product, nil
}

func (r *postgresProductRepository) GetProducts() ([]p.Product, error) {
	products := []p.Product{}

	query := "SELECT * FROM products"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product p.Product
		if err := rows.Scan(&product.Currency, &product.Description, &product.ID, &product.Name, &product.Price); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil

}
