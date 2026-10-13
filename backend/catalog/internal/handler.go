package internal

import (
	"fmt"
	"json"
	"net/http"

	db "github.com/RushilHKC/grpcSimulator/DB"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := db.GetProducts()
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "product id is: %v", id)
}
