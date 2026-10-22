package handlerSVC

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/service"
)

type Handler interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
	GetProduct(w http.ResponseWriter, r *http.Request)
}

type ProductHandler struct {
	Srv service.ProductSVC
}

var _ Handler = (*ProductHandler)(nil)

func NewHandler(srv service.ProductSVC) *ProductHandler {
	return &ProductHandler{Srv: srv}
}

func (h *ProductHandler) GetOrderDetails(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	details, err := h.Srv.GetOrderDetails(ctx, Id)

	if err != nil {
		log.Printf("Product not available: ", err)
	}

	json.NewEncoder(w).Encode(details)
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	products, err := h.Srv.GetProducts(ctx)

	if err != nil {
		log.Printf("Error getting Products: %v", err)
	}

	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	Id, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	product, err := h.Srv.GetProduct(ctx, Id)

	json.NewEncoder(w).Encode(product)

}
