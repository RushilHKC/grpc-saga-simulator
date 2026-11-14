package routes

import (
	"net/http"

	handlerSVC "github.com/RushilHKC/grpcSimulator/backend/catalog/internal/handler"
)

func Register(mux *http.ServeMux, handler handlerSVC.Handler) {

	mux.HandleFunc("GET /", handler.Welcome)
	mux.HandleFunc("GET  /products", handler.GetProducts)
	mux.HandleFunc("GET /product/{id}", handler.GetProduct)

}
