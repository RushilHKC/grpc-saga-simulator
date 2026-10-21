package service

import (
	"context"
	"time"

	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/env"
	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/model"
	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/repository"
)

type ProductSVC interface {
	GetProduct(ctx context.Context, id int) (*model.Product, error)
	GetProducts(ctx context.Context) ([]model.Product, error)
}

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) *ProductService {
	return &ProductService{repo: r}
}

func (srv *ProductService) GetProduct(ctx context.Context, id int) (*model.Product, error) {

	time_out := env.GetInt("database_time_out")
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time_out)*time.Second)
	defer cancel()

	return srv.repo.GetProduct(ctx, id)
}

func (srv *ProductService) GetProducts(ctx context.Context) ([]model.Product, error) {

	time_out := env.GetInt("database_time_out")

	ctx, cancel := context.WithTimeout(ctx, time.Duration(time_out)*time.Second)
	defer cancel()

	return srv.repo.GetProducts(ctx)

}
