package db

type product struct {
	id          int
	Name        string
	Description string
	Price       int
}

type DB struct {
	Products []product
}

// iphone := product{
// 	1,
// 	"Iphone17",
// 	"256GB lavender",
// 	88400,
// }

// television := product{
// 	1,
// 	"Iphone17",
// 	"256GB lavender",
// 	88400,
// }

// bike := product{
// 	1,
// 	"Iphone17",
// 	"256GB lavender",
// 	88400,
// }

// laptop := product{
// 	1,
// 	"Iphone17",
// 	"256GB lavender",
// 	88400,
// }

// smartwatch := product{
// 	1,
// 	"Iphone17",
// 	"256GB lavender",
// 	88400,
// }

// Products := []product{iphone, smartwatch, laptop, bike, television}

func GetProducts() []product {
	db := &DB{
		Products: []product{
			{1, "Iphone17", "256GB lavender", 88400},
			{1, "Iphone17", "256GB lavender", 88400},
			{1, "Iphone17", "256GB lavender", 88400},
			{1, "Iphone17", "256GB lavender", 88400},
			{1, "Iphone17", "256GB lavender", 88400},
			{1, "Iphone17", "256GB lavender", 88400},
			{1, "Iphone17", "256GB lavender", 88400},
		},
	}
	return db.Products
}
