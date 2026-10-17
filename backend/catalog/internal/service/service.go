package service

import (
	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/model"
	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/repository"
)	

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) *ProductService {
	return &ProductService{repo: r}
}

func (srv *ProductService) GetProduct(id int) (*model.Product, error) {
	return srv.repo.GetProduct(id)
}

func (srv *ProductService) GetProducts() ([]model.Product, error) {
	return srv.repo.GetProducts()
}
