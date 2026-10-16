package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := db.GetProducts()
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "product id is: %v", id)
}
