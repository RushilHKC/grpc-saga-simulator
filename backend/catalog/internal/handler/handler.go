package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/service"
)

type Handler struct {
	Srv *service.ProductService
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.Srv.GetProducts()

	if err != nil {
		log.Printf("Error getting Products: %v", err)
	}

	json.NewEncoder(w).Encode(products)
}

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	ID, err := strconv.Atoi(id)

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	product, err := h.Srv.GetProduct(ID)

	json.NewEncoder(w).Encode(product)
}
