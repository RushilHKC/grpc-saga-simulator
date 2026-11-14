package routes

import (
	"net/http"

	handlerSVC "github.com/RushilHKC/grpcSimulator/backend/catalog/internal/handler"
)

func NewMux(handler handlerSVC.Handler) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("GET  /products", handler.GetProducts)
	mux.HandleFunc("GET /product/{id}", handler.GetProduct)

	return mux
}
