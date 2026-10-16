// package db

// type product struct {
// 	id          int
// 	Name        string
// 	Description string
// 	Price       int
// }

// type DB struct {
// 	Products []product
// }

// // iphone := product{
// // 	1,
// // 	"Iphone17",
// // 	"256GB lavender",
// // 	88400,
// // }

// // television := product{
// // 	1,
// // 	"Iphone17",
// // 	"256GB lavender",
// // 	88400,
// // }

// // bike := product{
// // 	1,
// // 	"Iphone17",
// // 	"256GB lavender",
// // 	88400,
// // }

// // laptop := product{
// // 	1,
// // 	"Iphone17",
// // 	"256GB lavender",
// // 	88400,
// // }

// // smartwatch := product{
// // 	1,
// // 	"Iphone17",
// // 	"256GB lavender",
// // 	88400,
// // }

// // Products := []product{iphone, smartwatch, laptop, bike, television}

// func GetProducts() []product {
// 	db := &DB{
// 		Products: []product{
// 			{1, "Iphone17", "256GB lavender", 88400},
// 			{1, "Iphone17", "256GB lavender", 88400},
// 			{1, "Iphone17", "256GB lavender", 88400},
// 			{1, "Iphone17", "256GB lavender", 88400},
// 			{1, "Iphone17", "256GB lavender", 88400},
// 			{1, "Iphone17", "256GB lavender", 88400},
// 			{1, "Iphone17", "256GB lavender", 88400},
// 		},
// 	}
// 	return db.Products
// }

package main

import (
	"fmt"
	"log"

	repo "github.com/RushilHKC/grpcSimulator/backend/catalog/internal/Repository"
	config "github.com/RushilHKC/grpcSimulator/backend/catalog/internal/config"
	DB "github.com/RushilHKC/grpcSimulator/backend/catalog/internal/db"
	srv "github.com/RushilHKC/grpcSimulator/backend/catalog/internal/service"
)

func main() {

	con := config.Load_Db_Config()

	db, err := DB.New_db_connection(con.Port, con.User, con.Password, con.Db_name)

	if err != nil {
		log.Fatal("Error connecting DB")
	}
	productsRepo := repo.NewProductsRepository(db)
	productService := srv.NewProductService(productsRepo)

	fmt.Print(productService)

}
