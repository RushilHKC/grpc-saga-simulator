package service

import (
	db "github.com/RushilHKC/grpcSimulator/backend/catalog/internal/Repository"
)

type ProductService struct {
	// GetProductById(id int) Product
	// GetProducts() []Product
	// AddProduct(p Product)
	// RemoveProduct(id int)
	// UpdateProduct(id int)
	Repo *db.PostGressProductRepository
}

func NewProductService(r *db.PostGressProductRepository) *ProductService {
	return &ProductService{Repo: r}
}
