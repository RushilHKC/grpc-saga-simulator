package service

import (
	p "github.com/RushilHKC/grpcSimulator/backend/catalog/internal/model"
	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/repository"
)

type ProductService struct {
	// GetProductById(id int) Product
	// GetProducts() []Product
	// AddProduct(p Product)
	// RemoveProduct(id int)
	// UpdateProduct(id int)
	Repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) *ProductService {
	return &ProductService{Repo: r}
}

func (srv *ProductService) GetProduct(id int) (*p.Product, error) {
	return srv.Repo.GetProduct(id)
}

func (srv *ProductService) GetProducts() ([]p.Product, error) {
	return srv.Repo.GetProducts()
}
